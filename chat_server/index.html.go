/*
 * Copyright 2019 the go-netty project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

var indexHtml = []byte(`
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>WebSocket Chat</title>
</head>
<body>
	<script type="text/javascript">
		var socket;
		if (!window.WebSocket) {
			window.WebSocket = window.MozWebSocket;
		}
		if (window.WebSocket) {
			socket = new WebSocket("ws://127.0.0.1:8080/chat");
			socket.onmessage = function(event) {
				var cmd = JSON.parse(event.data);
				var ta = document.getElementById('responseText');
				ta.value = ta.value + '\n' + (cmd.name + ': ' + cmd.message);
			};
			socket.onopen = function(event) {
				var ta = document.getElementById('responseText');
				ta.value = "connection open!";
			};
			socket.onclose = function(event) {
				var ta = document.getElementById('responseText');
				ta.value = ta.value + "connection closed!";
			};
		} else {
			alert("Your browser does not support WebSocket!");
		}

		function send(name, message) {
			if (!window.WebSocket) {
				return;
			}
			if (socket.readyState == WebSocket.OPEN) {
				socket.send(JSON.stringify({"name" : name, "message" : message}));
			} else {
				alert("Connection is not open!");
			}
		}
	</script>
	<form onsubmit="return false;">
		<h3>WebSocket Chatroom:</h3>
		<textarea id="responseText" style="width: 500px; height: 300px;"></textarea>
		<br>
		<input type="text" name="name" style="width: 100px" value="Rob">
		<input type="text" name="message" style="width: 300px" value="Hello WebSocket">
		<input type="button" value="Send" onclick="send(this.form.name.value, this.form.message.value)">
		<input type="button" onclick="javascript:document.getElementById('responseText').value=''" value="Clear">
	</form>
	<br>
	<br>
</body>
</html>
`)
