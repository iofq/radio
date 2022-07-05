const title = document.getElementById("track")
const track_endpoint = "http://radio.10110110.xyz/current"

async function currentTrack() {
    const resp = await fetch(track_endpoint)
    const body = await resp.text()
    title.innerHTML = body
}
currentTrack()
window.setTimeout(currentTrack, 60 * 1000)
