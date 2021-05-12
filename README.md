# mssz-flyimg
This image is modified from the official image of flyimg.

An front page is added to the image, which can make it easier for users to achieve flyimg's features.

To run the image, please do remember to project the port 2222 and 80. And do rememb to provide your domain or IP with the port with -e DOMAIN="" when you run the image.

For axample, to quickly test the image:

docker run -it --rm -p 80:2222 -p 81:80 -e DOMAIN="xxxxxx:81" mssz/flyimg
