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
    });
  }
});
