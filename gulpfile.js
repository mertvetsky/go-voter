"use strict";

var gulp = require('gulp');
var connect = require('gulp-connect');
var open = require('gulp-open');
var browserify = require('browserify');
var reactify = require('reactify');
const babel = require('gulp-babel');
var source = require('vinyl-source-stream');
var lint = require('gulp-eslint');
var concat = require('gulp-concat');

var config = {
  paths: {
    html: './src/*.html',
    js: './src/**/*.js',
    images: './src/images/*',
    css: [
      'node_modules/bootstrap/dist/css/bootstrap.min.css',
      'node_modules/bootstrap/dist/css/bootstrap-theme.min.css'
    ],
    dist: './dist',
    mainJs: './src/main.js'
  }
};

//start a local devserver

gulp.task('html', function() {
  return gulp.src(config.paths.html)
      .pipe(gulp.dest(config.paths.dist))
      .pipe(connect.reload());
});

gulp.task('js', function() {
  return browserify(config.paths.mainJs)
    .transform('babelify', {presets: ["es2015", "react"]})
    .bundle()
    .on('error', console.error.bind(console))
    .pipe(source('bundle.js'))
    .pipe(gulp.dest(config.paths.dist + '/scripts'));
});

gulp.task('css', function() {
  return gulp.src(config.paths.css)
      .pipe(concat('bundle.css'))
      .pipe(gulp.dest(config.paths.dist + '/css'));
});

gulp.task('images', function() {
  return gulp.src(config.paths.images)
      .pipe(gulp.dest(config.paths.dist + '/images'));
});

gulp.task('lint', function() {
  return gulp.src(config.paths.js)
    .pipe(lint())
    .pipe(lint.format());
});

gulp.task('watch', function() {
  gulp.watch(config.paths.html, gulp.series('html'));
  gulp.watch(config.paths.mainJs, gulp.series('js', 'lint'));
});

gulp.task('default', gulp.parallel('html', 'js', 'css', 'images', 'lint', 'watch'));
