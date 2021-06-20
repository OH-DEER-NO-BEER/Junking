// mockMessage (VISIBLE!)
var mockP1 = {
	message: "Room Made",
	p1: {
		name: "wakishi",
		rate: {
			rock: 0.1,
			scissors: 0.1,
			paper: 0.8,
		},
	},
	p2: {
		name: "yokoro",
		rate: {
			rock: 0.3,
			scissors: 0.3,
			paper: 0.4,
		},
	},
};
var mockP2 = {
	message: "Room Entered",
	p1: {
		name: "wakishi",
		rate: {
			rock: 0.1,
			scissors: 0.1,
			paper: 0.8,
		},
	},
	p2: {
		name: "yokoro",
		rate: {
			rock: 0.3,
			scissors: 0.3,
			paper: 0.4,
		},
	},
};
var mockSelected;
// setting for data
let roomID = null;

//WebSocket
let ws = new WebSocket('wss://junking.tk/ws');

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
			console.log(mockSelected);
			clearInterval(intervalID);
		}
	}, 1000);
}
// for test
function mockSelectP1() {
	mockSelected = mockP1;
}
function mockSelectP2() {
	mockSelected = mockP2;
}
// do
clearStorage();
watchStorage();
