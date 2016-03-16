var gulp = require('gulp'),
  typescript = require('typescript'),
  ts = require('gulp-typescript'),
  sourcemaps = require('gulp-sourcemaps'),
  browserify = require('browserify'),
  source = require('vinyl-source-stream'),
  watch = require('gulp-watch'),
  del = require('del')
  ;

var project = ts.createProject('tsconfig.json', {typescript: typescript});

gulp.task('compile', function () {
  return gulp.src("src/**/*.{ts,tsx}")
    .pipe(sourcemaps.init())
    .pipe(ts(project))
    .js
    .pipe(sourcemaps.write())
    .pipe(gulp.dest('.'));
});

gulp.task('bundle', ['compile'], function () {
  var b = browserify('.tmp/tsBundle.js');
  return b.bundle()
    .pipe(source('bundle.js'))
    .pipe(gulp.dest('dist'))
    ;
});

gulp.task('clean', function (done) {
  del(['.tmp'], done.bind(this));
});

gulp.task('watch', function () {
  watch('src/**/*.{ts,tsx}', ['compile']);
});

gulp.task('default', ['watch']);
