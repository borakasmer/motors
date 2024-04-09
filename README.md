# motors
With this CLI Tool, Specific motors price list information is instantly retrieved by year. Motor Vehicle List information is printed on the screen by Parsing instantly from arabam.com. You can get different types of motors, with "-t 1, -t 2, -t 3" flags.<br><br>
&#x1F34E;<I>This is Arabam.com crawler CLI. In every request, It parses Arabam.com. If Arabam.com doesn't response, this service can't work.</I>
![CleanShot 2024-04-09 at 22 57 43](https://github.com/borakasmer/motors/assets/9459881/6e281b1a-6e49-46ce-8b23-5c8843829340) <br>

<b>Flags:</b>
<table><tr><td>🏍️:</td><td><I>motors</I></td></tr><tr><td></td><td><I>motors -t 2</T></td></tr></table>
<ul>
  <li> -t, --type int   Type of motor to be pulled from Arabam.com: '-t'(type of 4)(default type = 1 Cbr650R)</li>  
</ul>

<b>Default Command:</b> "motors" </br></br>
<b>Example Usage:</b>

<b>Usage:</b>
  motors [flags]
<ul>
  <li>"motors -t 3"</li>
  <li>"motors -t 2"</li> 
</ul>
**********************************************************************************************************</br>

<b><u>How to install Motors Cli:</u></b><br>

<table><tr><td><img src="https://user-images.githubusercontent.com/9459881/165053981-38543faf-4bae-4500-8c28-fd5f497e0f46.gif"></img></td>
  <td><i>go install github.com/borakasmer/motors@latest</i></td></tr></table>

<span style="color: red"><b>&#x1F534;Important:</b></span> You need Go program package to install Motors-Cli => <a href="https://go.dev/dl/" target="_blank">Go Downloads</a> </br>
<ol>
  <li>"go env" With the command "GOPATH" and "GOROOT" folders of GO can be seen.</li>
  <li>On <b>Mac</b> after "go install ..." the "go/bin/motors" file under GOPATH => should be copied under "go/bin/" folder under GOROOT.</li>
  <li>In <b>Windows</b>, there is no need to further action :white_check_mark:</li>
</ol>
<img width="427" alt="FRJfHssXEAAqNsP" src="https://user-images.githubusercontent.com/9459881/165074359-572ca085-b1bd-4dbc-840f-43b1690a6319.png">
<b>:green_book: Extra Detail</b><br>
--------------------------------------------------------------------------------------------------------------
<ul>
  <li> "-t" Type of Motors in Specific list. The default value is 1. List : 1-)Cbr650r 2-)Zx6R 3-)Cbr600RR 4-)R25 </li>  
  <li> "motors" command => By default it means "motors --type 1".</li>
  <li> "More Motors Will be Added Soon......".</li>
<ul>
    

