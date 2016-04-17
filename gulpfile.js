var gulp = require('gulp'),
  $ = require('gulp-load-plugins')(),
  browserify = require('browserify'),
  watchify = require('watchify'),
  source = require('vinyl-source-stream'),
  buffer = require('vinyl-buffer'),
  glob = require('glob'),
  runSequence = require('run-sequence'),
  del = require('del');

var debug = true;

var paths = {
  src: "./src/**/*.{ts,tsx,js}",
  style: "./src/**/*.{css,sass,scss}",
  out: 'dist'
};

var files = {
  src: glob.sync(paths.src),
  style: glob.sync(paths.style),
};

function buildSrc(files, watch) {
  var props = watchify.args;
  props.entries = files;
  props.debug = debug;

  var bundler = browserify(props).plugin('tsify');
  if (watch) {
    bundler = watchify(bundler);
  }
  function rebundle() {
    return bundler.bundle()
      .on("error", function () {
        return $.notify.onError({
          title: "Bundle error",
          message: "<%= error.message %>"
        });
      })
      .pipe(source('bundle.js'))
      .pipe(buffer())
      .pipe($.sourcemaps.init({loadMaps: true}))
      //.pipe($.uglify())
      .pipe($.sourcemaps.write({
        includeContent: false,
        sourceRoot: '..'
      }))
      .pipe(gulp.dest(paths.out));
  }

  bundler.on('update', function () {
    rebundle();
    $.util.log("Bundled.");
    $.util.log(paths.src);
    $.notify("Scripts bundled");
  });
  bundler.on('log', function(message) {
    console.log(message);
  });
  return rebundle();
}

gulp.task("build-script", function () {
  return buildSrc(files.src, false);
});

gulp.task("build-styles", ["sass"]);

gulp.task("sass", function () {
  return gulp.src(paths.style)
    .pipe($.sourcemaps.init())
    .pipe($.sass())
    .on("error", $.sass.logError)
    .pipe($.concatCss("main.css"))
    .pipe($.sourcemaps.write())
    .pipe(gulp.dest('./dist'))
    .pipe($.notify("Styles bundled"));
});

gulp.task("watch-script", function () {
  return buildSrc(files.src, true);
});

gulp.task("default", function (done) {
  runSequence(["build-script", "build-styles"], done);
});

gulp.task("watch", ["watch-script"], function() {
  return $.watch(paths.style, ["build-styles"]);
});

