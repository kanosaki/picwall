var gulp = require('gulp'),
  $ = require('gulp-load-plugins')(),
  browserify = require('browserify'),
  watchify = require('watchify'),
  source = require('vinyl-source-stream'),
  glob = require('glob'),
  del = require('del');

var project = $.typescript.createProject('tsconfig.json', {typescript: require('typescript')});

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
    var stream = bundler.bundle();
    stream
      .on("error", function () {
        $.notify.onError({
          title: "Bundle error",
          message: "<%= error.message %>"
        })
      })
      .pipe(source('bundle.js'))
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

