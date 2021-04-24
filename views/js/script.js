$(window).on("scroll", function () {
    stop = Math.round($(window).scrollTop());
    if (stop > $("header").height()) {
        $("header").css({
            position: "fixed",
            "background-color": "#2e2e2e"
        });
    } else {
        $("header").css({
            position: "absolute",
            "background-color": "transparent",
        });
    }
});
