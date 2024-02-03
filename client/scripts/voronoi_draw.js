export function drawVoronoiDiagramOnCanvas(canvas, points) {
  const voronoi = new Voronoi();
  const bbox = { xl: 0, xr: canvas.width, yt: 0, yb: canvas.height };
console.log(points);
  const diagram = voronoi.compute(points, bbox);
	console.log(diagram);

  const ctx = canvas.getContext('2d');
  ctx.clearRect(0, 0, canvas.width, canvas.height);


  points.forEach(point => {
	ctx.beginPath();
	ctx.arc(point.x, point.y, 1, 0, 2 * Math.PI);
	ctx.fillStyle = "#0000ff"; // Blue color for points
	ctx.fill();
	// ctx.stroke();
  });

  for (const cell of diagram.cells) {
    ctx.beginPath();
    ctx.lineWidth = 1; // Line width
    ctx.moveTo(cell.halfedges[0].getStartpoint().x, cell.halfedges[0].getStartpoint().y);


	

    for (const halfEdge of cell.halfedges) {
	  const endpoint = halfEdge.getEndpoint();
		console.log(endpoint.x, endpoint.y);
      ctx.lineTo(endpoint.x, endpoint.y);
      // ctx.lineTo(halfEdge.endpoint.x, halfEdge.endpoint.y);
    }

    ctx.closePath();
    ctx.stroke();
  }
}
