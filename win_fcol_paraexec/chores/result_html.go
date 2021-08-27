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

package chores

const CmdResTemplate = `
<!--
 Copyright (C) 2021 kmahyyg
 
 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU Affero General Public License as
 published by the Free Software Foundation, either version 3 of the
 License, or (at your option) any later version.
 
 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU Affero General Public License for more details.
 
 You should have received a copy of the GNU Affero General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.
-->
<!DOCTYPE html>
<html>
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <meta http-equiv="X-UA-Compatible" content="ie=edge">
        <meta name="description" content="">
        <title>Result of fCol</title>
        <!-- introduce bootstrap -->
        <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.0/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-KyZXEAg3QhqLMpG8r+8fhAXLRk2vvoC2f3B09zVXn8CA5QIVfZOJ3BCsw2P0p/We" crossorigin="anonymous">
        <link rel="stylesheet" href="https://cdn.jsdelivr.net/gh/highlightjs/cdn-release@11.2.0/build/styles/default.min.css" crossorigin="anonymous">
        <style>
            .mytitle{
                vertical-align: middle;
                text-align: center;
                margin: 0px auto;
            }
            #footer{
                display: table;
                text-align: center;
                margin-left: auto;
                margin-right: auto;
            }
        </style>
    </head>
    <body>
        <div class="container align-middle">
            <div class="row"><h1 class="mytitle">Scan Result of fCol</h1></div>
            <div class="row align-middle">
                <div class="col-3">
                    Comment 指令用途与备注
                </div>
                <div class="col-3">
                    Command 指令
                </div>
                <div class="col-6">
                    Execution Result 指令结果
                </div>
            </div>
            {{range .CmdOutputs}}
            <div class="row align-middle">
                <div class="col-3">
                    {{.CmdComment}}
                </div>
                <div class="col-3">
                    {{.CmdDetail}}
                </div>
                <div class="col-6">
                    <pre><code>{{.CmdOutput}}</code></pre>
                </div>
            </div>
            {{end}}
        </div>
    </body>
    <!-- introduce bootstrap -->
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.1.0/dist/js/bootstrap.bundle.min.js" integrity="sha384-U1DAWAznBHeqEIlVSCgzq+c9gqGAJn5c/t99JyeKa9xxaYpSvHU5awsuZVVFIhvj" crossorigin="anonymous"></script>
    <script src="https://cdn.jsdelivr.net/gh/highlightjs/cdn-release@11.2.0/build/highlight.min.js" crossorigin="anonymous"></script>
    <script>hljs.highlightAll();</script>
    <br>
    <footer>Copyright (C) 2021 Patrick Young - Project Ha1l </footer>
</html>
`
