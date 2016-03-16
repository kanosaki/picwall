var gulp = require('gulp'),
  $ = require('gulp-load-plugins')(),
  browserify = require('browserify'),
  watchify = require('watchify'),
  source = require('vinyl-source-stream'),
  buffer = require('vinyl-buffer'),
  glob = require('glob'),
  del = require('del');

var paths = {
  srcFiles: glob.sync("./src/**/*.{ts,tsx,js}"),
  out: 'dist'
};

function buildSrc(files, watch) {
  var props = watchify.args;
  props.entries = files;
  props.debug = true;

  var bundler = browserify(props).plugin('tsify');
  if (watch) {
    bundler = watchify(bundler);
  }
  function rebundle() {
    return bundler.bundle()
      .on("error", function () {
        $.notify.onError({
          title: "Bundle error",
          message: "<%= error.message %>"
        })
      })
      .pipe(source('bundle.js'))
      .pipe(buffer())
      .pipe($.sourcemaps.init({loadMaps: true}))
      .pipe($.uglify())
      .pipe($.sourcemaps.write({
        includeContent: false,
        sourceRoot: '..'
      }))
      .pipe(gulp.dest(paths.out));
  }

  bundler.on('update', function () {
    rebundle();
    $.util.log("Bundled.");
    $.util.log(paths.srcFiles);
  });
  return rebundle();
}

gulp.task("build", function (done) {
  buildSrc(paths.srcFiles, false);
  done();
});

gulp.task("default", ["build"]);

gulp.task("watch", ["default"], function () {
  return buildSrc(paths.srcFiles, true);
});

