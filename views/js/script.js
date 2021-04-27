$(window).on("scroll", function () {
    stop = Math.round($(window).scrollTop());
    if (stop > $("header").height()) {
        $("header").css({
            position: "fixed",
            "background-color": "#121212"
        });
    } else {
        $("header").css({
            position: "absolute",
            "background-color": "transparent",
        });
    }
});
