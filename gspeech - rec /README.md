# gspeech-rec

### Google's Speech Recognition from the command line

**speech-rec.sh** is a bash script that sends a flac audio file to Google for speech recognition and prints the N-best returned hypotheses. It uses the (until now) freely available, though up to 50 requests a day limited, Google ASR API version 2.

More interestingly, if you have `sox` (or `parecord` and `timeout`) installed on your machine, you can use the script to interactively record and utterance and get its transcription.

To install sox type:

    sudo apt-get install sox

For more details about Google's ASR service check out [this article](https://aminesehili.wordpress.com/2015/02/08/on-the-use-of-googles-speech-recognition-api-version-2/).

# Some examples

### Recognize speech from audio file
    ./speech-rec.sh -i record.flac --rate 16000

### Try other languages (default=en_US):

Use option `-l` to set language:

**French**

    ./speech-rec.sh -i record.flac --rate 16000 --language fr_FR

**Spanish**

    ./speech-rec.sh -i record.flac -r 16000 -l es_ES

**German**

    ./speech-rec.sh -i record.flac -r 16000 -l de_DE

### Simply say something and get what you've said

This will give you 3 seconds to talk:

    ./speech-rec.sh
   
### Talk for 7 seconds (default=3)

Use option `-d` to set duration:

    ./speech-rec.sh -d 7
   
The script writes audio data into an audio file named `record_DATE_TIME.flac` where `DATE_TIME` is the output of:

    date "+%Y%b%d_%H-%M-%S"
    
To replay your utterance with `sox` type:

    play record.flac
    
or with `paplay`:

    paplay record.flac
   
# Requirements
### Google API Key
The key delivered with this script is to be used for test purposes only. It may be disabled by Google anytime. You should generate and use your own API key (follow this link for key generation instructions: http://www.chromium.org/developers/how-tos/api-keys)

#### `sox`
for audio data acquistion, or

#### `parecord` + `timeout`

if you don't have `sox`.

#### `wget`
to send data to server.


# Python API

[this repository](https://github.com/Uberi/speech_recognition/) offers a Python ASR API for many vendors.

# Author
Amine SEHILI (<amine.sehili@gmail.com>)

# License
Copyright (C) Amine SEHILI 2015-2016.

This program is available under the GNU GENERAL PUBLIC LICENSE Version 3.

Source code on GitHub: https://github.com/amsehili/gspeech-rec.


