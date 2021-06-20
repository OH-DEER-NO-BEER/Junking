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
var mockP1Result = {
	P1Hand: "rock",
	P2Hand: "scissors",
	result: "win",
};
var mockP2Result = {
	P1Hand: "rock",
	P2Hand: "scissors",
	result: "lose",
};
var mockSelected;
var mockSelectedResult;

// setting for data
let roomID = null;
//WebSocket
// let ws = new WebSocket('wss://junking.tk:8080/ws/');

// methods related to storage
function sentToUnity() {}
function getRoomID() {
	return localStorage.getItem("roomID");
}
function getSelectHand() {
	return localStorage.getItem("selectHand");
}
function clearStorage() {
	localStorage.clear();
}
function watchStorageRoomID() {
	let intervalID = window.setInterval(function() {
		roomID = getRoomID();
		if (roomID != null) {
			// if roomID is written by Unity, go out of block
			uniIns.SendMessage(
				// Send to Unity
				"EventControl",
				"SetJson",
				JSON.stringify(mockSelected)
			);
			clearInterval(intervalID);
		}
	}, 1000);
}
function watchStorageSelectHand() {
	let intervalID = window.setInterval(function() {
		selectHand = getSelectHand();
		if (selectHand != null) {
			console.log(mockSelectedResult);
			uniIns.SendMessage(
				// Send to Unity
				"EventControl",
				"SetJson",
				JSON.stringify(mockSelectedResult)
			);
			clearInterval(intervalID);
		}
	}, 1000);
}
// for test
function mockSelectP1() {
	mockSelected = mockP1;
	mockSelectedResult = mockP1Result;
}
function mockSelectP2() {
	mockSelected = mockP2;
	mockSelectedResult = mockP2Result;
}
// do
clearStorage();
watchStorageRoomID();
watchStorageSelectHand();
