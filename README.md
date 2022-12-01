# bjj-tools

A repo for that holds a collection of scripts written in Golang made
for enhancing bjjfanatics video formats. The way their app / website is
setup is that you have a number of timestamps in each video which
correspond with different techniques. Purchased videos can be downloaded
but these videos do not come with the timestamps so watching offline or
on desktop the experience is a lot worse compared to their mobile app
or website. Ontop of that rarely are instructionals made for a person to
sit through them in a single sitting, instead a practitioner will watch
a limited amount of techniques and drill accordingly or make notes.

`.mp4` files have chapter functionality which allow the user to jump
between points in the file quickly. These scripts are designed to help
the user utilize this to have a better learning experience.


## Video Merge<https://github.com/heyjp/bjj-tools/video_merge>

Utilizes the `extractor` module alongside ffmpeg to extract the metadata
from the given video file location, updates the metadata to include
timestamps from the given chapters folder and then combines the updated
metadata file with chapters and merges with the video file to create a
brand new video.

Merges chapters, videos and metadata together to place chapters inside
of videos.


# Requirements

For using Video merge:

1) Download `ffmpeg` [here](https://ffmpeg.org/download.html)
2) (Optional): Add ffmpeg directory to your computers environment variables
3) Place the `video-merge.exe` alongside your `.mp4` files.
4) Move the folder holding your chapters files to the same directory
5) Run `video-merge.exe`


## Other Scripts

### Fanatics Search<https://github.com/heyjp/bjj-tools/fanatics_search>

When passed in a product name `power-ride-by-craig-jones` on bjjfanatics
and a download location will extract timestamps from the bjjfanatics
website and convert them into a text file named `chapters-<n>.txt`


### Fanatics Crawler<https://github.com/heyjp/bjj-tools/fanatics_crawler>

When run crawls the fanatics website and returns a document with the
links to the products as well as the folder you wish save the product
timestamps in.


### Fanatics Chapters<https://github.com/heyjp/bjj-tools/fanatics_search>

Takes the file created by the fanatics Crawler and retrieves the
timestamps for each product and places them into the designated folder.


### Misc

- Metadata - Prepares the metadata file, handles clearing old chapters
  if they exist and adding new ones.
- Chapters - Prepares the chapters file/s after receiving data from the
  web search
- Dircheck - Tools for getting directory / file data.
- Extractor - Tool for using ffmpeg

Prepares
