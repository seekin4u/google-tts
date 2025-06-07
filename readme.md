https://github.com/ExistentialAudio/BlackHole?tab=readme-ov-file#installation-instructions

mpg123 should be on macos by default

create credentials and download them
enable api on https://console.cloud.google.com/apis/api/texttospeech.googleapis.com/
and export GOOGLE_APPLICATION_CREDENTIALS=~/Documents/google-auth.json

https://cloud.google.com/text-to-speech/pricing?hl=en
this app will use free tier voices with up to 4mil characters/mo

settings are shown in ./images
Default system input and output should be chosen in system settings and then in discord's voice settings

go build .
./google-tts "текст для озвучки"
