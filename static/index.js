const title = document.getElementById("track")
const track_endpoint = "https://radio.10110110.xyz/current"

const count = document.getElementById("count")
const count_endpoint = "https://radio.10110110.xyz/listeners"

const audio = document.getElementById("audio")
const skip_endpoint = "https://radio.10110110.xyz/skip"

async function connListeners() {
    const resp = await fetch(count_endpoint)
    const body = await resp.text()
    count.innerHTML = body
}

async function currentTrack() {
    const resp = await fetch(track_endpoint)
    const body = await resp.text()
    title.innerHTML = body
}

async function skipTrack() {
    await fetch(skip_endpoint)
    currentTrack()
    audio.pause();
    audio.play();
}

audio.volume = 0.5;
currentTrack()
connListeners()
window.setInterval(currentTrack, 5 * 1000)
window.setInterval(connListeners, 5 * 1000)
