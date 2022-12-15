# bjj-tools

A repo for that holds a collection of scripts written in Golang made
for enhancing bjjfanatics video formats. The way their app / website is
setup is that you have a number of timestamps in each video. Purchased
videos can be downloaded but these videos do not come with timestamps so
watching offline or on desktop is a worse experience.

`.mp4` files have chapter functionality which allow the user to jump
between points in the video quickly. These scripts are designed to
create / merge timestamps with their related videos to help the user have a better learning experience.


### [Video
Merge](https://github.com/Heyjp/bjj-tools/tree/main/video_merge)

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
2) Place the `video-merge.exe` alongside your `.mp4` files.
4) Move the folder holding your chapters files to the same directory
5) Run `video-merge.exe`


### [Chapters](https://github.com/heyjp/bjj-tools/tree/main/chapters)

A collection of scripts for collecting timestamps.

Usage: 
`$ main <crawl|chapters|search>`


#### Crawl

Command will crawl through the site and will create a text file with two
values. `<product name> <location for chapters file>`. You can choose
between crawling the entire site and retrieving every product or
crawling the first couple of pages  of their new products lineup. 

`$ main crawl <all|new>`

If a product name exists already it will not be written over.


#### chapters

If a `fanatics-products.txt` file has been created the `chapters` option
will then create a `chapters/` folder for each listed product.

`$ main chapters`


#### search

Like chapters, creates a chapters folder with timestamps but targets a
specific product entered by the user. 

`$ main search <product url>`


### Bugs

If you find an issue or instance where the scripts do not work or would
like to suggest a feature feel free to create an issue.
