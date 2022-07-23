let end = Date.now() + 15 * 200;

// go Buckeyes!
let colors = ["#bb0000", "#a864fd", "#29cdff", "#78ff44", "#ff718d", "#fdff6a"];

(function frame() {
  confetti({
    particleCount: 6,
    angle: 60,
    spread: 55,
    origin: { x: 0 },
    colors: colors,
  });
  confetti({
    particleCount: 6,
    angle: 120,
    spread: 55,
    origin: { x: 1 },
    colors: colors,
  });

  if (Date.now() < end) {
    requestAnimationFrame(frame);
  }
})();
