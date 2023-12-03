export function drawLine() {
    // Get the canvas element and its context
    var canvas = document.getElementById("mainCanvas");
    var ctx = canvas.getContext("2d");

    // Draw a line on the canvas
    ctx.beginPath();
    ctx.moveTo(50, 50);  // Starting point (x, y)
    ctx.lineTo(350, 150); // Ending point (x, y)
    ctx.strokeStyle = "#ff0000"; // Red color
    ctx.lineWidth = 5; // Line width
    ctx.stroke(); // Draw the line

	// Draw a sample point 50, 50
	ctx.beginPath();
	ctx.arc(39, 260, 5, 0, 2 * Math.PI);
	ctx.fillStyle = "#0000ff"; // Blue color for points
	ctx.fill();
	ctx.stroke();
	

	// Function to fetch points and draw them on the canvas
    function fetchAndDrawPoints() {
        var leftCoordinate = document.getElementById("left_coordinate").value;
        var rightCoordinate = document.getElementById("right_coordinate").value;

        // Make an HTTP GET request to the server
        fetch(`http://127.0.0.1:8080/generatePoints?left=${leftCoordinate}&right=${rightCoordinate}`)
            .then(response => response.json())
            .then(points => {
                // Clear the canvas
                ctx.clearRect(0, 0, canvas.width, canvas.height);

                // Draw the points on the canvas
                points.forEach(point => {
                    ctx.beginPath();
                    ctx.arc(point.x, point.y, 1, 0, 2 * Math.PI);
                    ctx.fillStyle = "#0000ff"; // Blue color for points
                    ctx.fill();
                    // ctx.stroke();
                });
            })
            .catch(error => {
                console.error('Error fetching points:', error);
            });
    }

    // Add event listener for the button click
    document.getElementById("submit_coordinate").addEventListener("click", fetchAndDrawPoints);
}
