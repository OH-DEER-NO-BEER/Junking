// setting for data
let roomID = null;
let data = {
	roomID: null,
};
// methods related to storage
function sendJson() {
	console.log("	");
}
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
			console.log("Caught!!!!!!!!!!!!!");
			clearInterval(intervalID);
		}
	}, 1000);
}
// do
clearStorage();
watchStorage();
