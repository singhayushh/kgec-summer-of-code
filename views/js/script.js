$('.js--scroll-to-contact').click(function () {
  $('html, body').animate(
    { scrollTop: $('#contact').offset().top - 150 },
    2000
  );
});

$('.js--scroll-to-service').click(function () {
  $('html, body').animate(
    { scrollTop: $('#about .flex-1').offset().top - 100 },
    1500
  );
});

$('.hero-btn').click(function () {
  $('html, body').animate({ scrollTop: $('#about').offset().top }, 1500);
});

$(window).on('scroll', function () {
  stop = Math.round($(window).scrollTop());
  if (stop > $('header').height()) {
    $('nav').css({
      position: 'fixed',
      'background-color': '#fff',
      'z-index': '999',
      'box-shadow': 'rgba(0, 0, 0, 0.1) 6px 6px 13px',
    });
    $('.logo img').css({
      filter: 'invert(0)',
    });
    $(
      '.menu li a, .menu li a:hover, .menu li a:active, .menu li a:visited'
    ).css({
      color: '#0d1328',
    });
  } else {
    $('nav').css({
      position: 'absolute',
      'background-color': 'transparent',
      'box-shadow': 'none',
    });
    $('.logo img').css({
      filter: 'invert(1)',
    });
    $(
      '.menu li a, .menu li a:hover, .menu li a:active, .menu li a:visited'
    ).css({
      color: '#e2e1e4',
    });
  }
});

/* Add Animations on scroll */
$('.js--wp-1').waypoint(
  function (direction) {
    if (direction === 'down') {
      $(this.element).addClass('animated fadeInUp');
      this.destroy();
    }
  },
  {
    offset: '80%',
  }
);
$('.js--wp-2').waypoint(
  function (direction) {
    if (direction === 'down') {
      $(this.element).addClass('animated fadeIn');
      this.destroy();
    }
  },
  {
    offset: '80%',
  }
);
$('.js--wp-3').waypoint(
  function (direction) {
    if (direction === 'down') {
      $(this.element).addClass('animated fadeInRight');
      this.destroy();
    }
  },
  {
    offset: '80%',
  }
);
$('.js--wp-4').waypoint(
  function (direction) {
    if (direction === 'down') {
      $(this.element).addClass('animated flipInX');
      this.destroy();
    }
  },
  {
    offset: '80%',
  }
);
$('.js--wp-5').waypoint(
  function (direction) {
    if (direction === 'down') {
      $(this.element).addClass('animated fadeInLeft');
      this.destroy();
    }
  },
  {
    offset: '80%',
  }
);
$('.js--wp-6').waypoint(
  function (direction) {
    if (direction === 'down') {
      $(this.element).addClass('animated fadeInDown');
      this.destroy();
    }
  },
  {
    offset: '80%',
  }
);
const options = {
  bottom: '64px', // default: '32px'
  right: 'unset', // default: '32px'
  left: '32px', // default: 'unset'
  time: '0.5s', // default: '0.3s'
  mixColor: '#fff', // default: '#fff'
  backgroundColor: '#fff',  // default: '#fff'
  buttonColorDark: '#100f2c',  // default: '#100f2c'
  buttonColorLight: '#fff', // default: '#fff'
  saveInCookies: true, // default: true,
  label: '', // default: ''
  autoMatchOsTheme: true // default: true
}

const darkmode = new Darkmode(options);
darkmode.showWidget();
