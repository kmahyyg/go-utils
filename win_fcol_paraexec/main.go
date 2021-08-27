// Copyright (C) 2021 kmahyyg
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"bufio"
	"fcol_paraexec/chores"
	"log"
	"os"
	"os/exec"
	"runtime"
	"sync"
	"text/template"
)

func execCmd(command string) string {
	cmd := exec.Command("C:\\Windows\\System32\\cmd.exe", "/c", command)
	outputs, err := cmd.CombinedOutput()
	var foutputs []byte
	if err != nil {
		foutputs = []byte("Fatal Error: " + err.Error() + "\r\n\r\n")
	}
	foutputs = append(foutputs, outputs...)
	return string(foutputs)
}

func parseAndExec(sIptChan chan *chores.CmdOutput, resLst *chores.FinalResults, wg *sync.WaitGroup) {
	log.Println("Debug: Start Processing Task")
	for {
		tempCmd, hasMore := <-sIptChan
		if hasMore {
			tempCmd.CmdOutput = execCmd(tempCmd.CmdDetail)
			resLst.Add(tempCmd)
			log.Println("Debug: End Processing Task")
		} else {
			wg.Done()
			break
		}
	}
}

func main() {
	runThreads := runtime.NumCPU() - 1
	runtime.GOMAXPROCS(runThreads)
	log.Println("Start Working...")
	// initiate storage
	var finalres = &chores.FinalResults{
		CmdOutputs: make([]*chores.CmdOutput, 0),
		Mu:         &sync.RWMutex{},
	}
	iptchan := make(chan *chores.CmdOutput)
	// initiate html template engine
	tpl, err := template.New("cmdRes").Parse(chores.CmdResTemplate)
	if err != nil {
		log.Fatal(err)
	}
	// check data validity
	if len(chores.Exec_Collect_Cmds) != len(chores.Exec_Collect_Cmds_Comment) {
		log.Fatal("Commands are not corresponding to its comments.")
	}
	// build coroutine wait group
	var wg = sync.WaitGroup{}
	// build workers and go for tasks
	for i := 0; i < runThreads; i++ {
		log.Println("Debug: Start Build worker, i=", i)
		wg.Add(1)
		go parseAndExec(iptchan, finalres, &wg)
		log.Println("Debug: End Build worker, i=", i)
	}
	// parse data and go
	j := 0
	for j < len(chores.Exec_Collect_Cmds)-1 {
		log.Println("Debug: Start Send Tasks to Input Channel, j=", j)
		var tempCmd = &chores.CmdOutput{
			CmdDetail:  chores.Exec_Collect_Cmds[j],
			CmdComment: chores.Exec_Collect_Cmds_Comment[j],
		}
		iptchan <- tempCmd
		j++
		log.Println("Debug: End Send Tasks to Input Channel, j=", j)
	}
	close(iptchan)
	// read the output and close the channels
	wg.Wait()
	// generate html
	f, err := os.Create("resultb.html")
	if err != nil {
		log.Fatalln(err)
	}
	html_w := bufio.NewWriter(f)
	err = tpl.Execute(html_w, finalres)
	if err != nil {
		log.Fatalln(err)
	}
	// flush buffer and close file
	err = html_w.Flush()
	if err != nil {
		log.Fatalln(err)
	}
	err = f.Close()
	if err != nil {
		log.Fatalln(err)
	}
	// log exit
	log.Println("End Working...")
}
