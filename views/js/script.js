$(window).on("scroll", function () {
  stop = Math.round($(window).scrollTop());
  if (stop > $("header").height()) {
    $("header").css({
      position: "fixed",
      "background-color": "rgba(12, 32, 43, 0.503)",
      "backdrop-filter": "blur(10px)",
    });
  } else {
    $("header").css({
      position: "absolute",
      "background-color": "transparent",
      "backdrop-filter": "blur(0px)",
    });
  }
});

// anime({
//   targets: "#ksocTitle path",
//   strokeDashoffset: [anime.setDashoffset, 0],
//   easing: "easeInOutSine",
//   duration: 2500,
//   delay: function (el, i) {
//     return i * 250;
//   },
//   direction: "alternate",
//   loop: false,
// });
