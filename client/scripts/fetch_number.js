export function fetchNumberOnClick() {
	return
    // Add event listener for form submission
    document.getElementById("submit_coordinate").addEventListener("click", function() {
        // Get left and right coordinates from the form
        var leftCoordinate = document.getElementById("left_coordinate").value;
        var rightCoordinate = document.getElementById("right_coordinate").value;

        // Call fetchNumber with the coordinates as parameters
        fetchNumber(leftCoordinate, rightCoordinate);
    });
}

export function fetchNumber(leftCoordinate, rightCoordinate) {
	// Get the <p> element with id "message"
    var messageElement = document.getElementById("message");

    // Make an HTTP GET request to a public domain (example: JSONPlaceholder)
    fetch(`https://jsonplaceholder.typicode.com/todos/1?left=${leftCoordinate}&right=${rightCoordinate}`)
        .then(response => response.json())
        .then(data => {
            // Extract a number from the fetched data (assuming the response contains a number)
            var fetchedNumber = data.id;

            // Update the <p> element with the fetched number
            messageElement.textContent = 'Fetched Number: ' + fetchedNumber;
        })
        .catch(error => {
            console.error('Error fetching data:', error);
        });
}
