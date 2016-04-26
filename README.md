Upload
======

Golang file uploader


Use the program to upload files through html form.
To test the service:

curl -F file=@local-file http://localhost:4000/upload/

The file will now be written to the same directory as the program is in.
The permission is set to 700 and the name of the file is the same as it was when it was sent.

//Christopher 21/10-13
