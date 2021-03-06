/** 
 * gulp依赖
 * npm install  --save-dev gulp gulp-minify-css gulp-autoprefixer gulp-sourcemaps gulp-jshint gulp-uglify gulp-rename gulp-concat gulp-clean gulp-webpack
 */

var gulp = require('gulp'),
    minifycss = require('gulp-minify-css'),
    autoprefixer = require('gulp-autoprefixer'),
    jshint = require('gulp-jshint'),
    uglify = require('gulp-uglify'),
    rename = require('gulp-rename'),
    sourcemaps = require('gulp-sourcemaps'),
    gulpwebpack = require('gulp-webpack'),
    webpack = require('webpack'),
    concat = require('gulp-concat'),
    clean = require('gulp-clean');

var notify = require("gulp-notify");
var markdown = require('gulp-markdown');
 

/**
 * 删除css
 */
gulp.task('clean-css', function(){
    return gulp.src('./public/css/app.min.css')
        .pipe(clean());
});


/**
 * 合并css
 */
gulp.task('css', ['clean-css'], function(){
    return gulp.src('./public/css/app.css')
        .pipe(sourcemaps.init())
        .pipe(minifycss({rebase: false, advanced: false, keepBreaks: true}))
        .pipe(autoprefixer())
        .pipe(sourcemaps.write())
        .pipe(rename('app.min.css'))
        .pipe(gulp.dest('./public/css/'));
});

/**
 * 删除js
 */
gulp.task('clean-js', function(){
    return gulp.src('./public/js/+(app|login|lib)?(.min).js')
        .pipe(clean());
});

/**
 * webpack 编译
 */
gulp.task('webpack', ['clean-js'], function(cb){
    return gulp.src('./app/app/+(app|login|lib).js')
        .pipe(gulpwebpack( require('./webpack.config.js') ))
        .pipe(gulp.dest('./public/js/'))
});

/**
 * 压缩js
 */
gulp.task('js', ['webpack'], function(){
    return gulp.src('./public/js/+(app|lib|login).js')
        // .pipe(jshint())
        // .pipe(uglify())
        .pipe(rename({suffix: ".min"}))
        .pipe(gulp.dest('./public/js/'))
        .pipe(notify('<%= file.relative %> build done!'));
});

gulp.task("watch", ['default'], function(){
    gulp.watch('./**/*.vue', ['js']);
    gulp.watch('./public/components/**/**.js', ['js']);
    gulp.watch('./app/app/**.js', ['js']);
    // gulp.watch('./public/css/app.css', ['css']);
    gulp.watch('./public/css/main.css', ['css']);
});

gulp.task('default', ['css', 'js']);

gulp.task('uglify', ['default'], function(){
    return gulp.src('./public/js/+(app|lib|login).js')
        .pipe(uglify())
        .pipe(rename({suffix: ".min"}))
        .pipe(gulp.dest('./public/js/'));
})

gulp.task('md', function () {
    return gulp.src('release-note.md')
        .pipe(markdown())
        .pipe(rename('release-note.html'))
        .pipe(gulp.dest('./'));
});