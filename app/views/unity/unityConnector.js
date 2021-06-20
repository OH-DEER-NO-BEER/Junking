// mockMessage (VISIBLE!)
var mockMessageScene2 = {
	message: "from JavaScript!",
	myself: {
		name: "wakishi",
		rate: {
			rock: 0.1,
			scissors: 0.1,
			paper: 0.8,
		},
	},
	opponent: {
		name: "yokoro",
		rate: {
			rock: 0.3,
			scissors: 0.3,
			paper: 0.4,
		},
	},
};
var mockMessageScene3 = {
	message: "from JavaScript!",
	myself: {
		name: "wakishi",
		rate: {
			rock: 0.1,
			scissors: 0.1,
			paper: 0.8,
		},
	},
	opponent: {
		name: "yokoro",
		rate: {
			rock: 0.3,
			scissors: 0.3,
			paper: 0.4,
		},
	},
};
// setting for data
let roomID = null;

//WebSocket
let ws = new WebSocket('ws://' + window.location.host + '/ws/');

// methods related to storage
function sentToUnity() {}
function getRoomID() {
	return localStorage.getItem("roomID");
}
function clearStorage() {
	localStorage.clear();
}
function watchStorage() {
	let intervalID = window.setInterval(function() {
		roomID = getRoomID();
		console.log("now: " + roomID);
		if (roomID != null) {
			// if roomID is written by Unity, go out of block
			ws.addEventListener('message', function(e) {
				console.log(e);
				uniIns.SendMessage( // SEND to Unity
					"EventControl",
					"SetJson",
					JSON.stringify(mockMessageScene2)
				);				
			});
			ws.send(
				JSON.stringify({
					"roomID": roomID
				})
			);
			clearInterval(intervalID);
		}
	}, 1000);
}



// do
clearStorage();
watchStorage();
