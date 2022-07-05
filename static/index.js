const title = document.getElementById("track")
const track_endpoint = "http://localhost:8000/current"

async function currentTrack() {
    const resp = await fetch(track_endpoint)
    const body = await resp.text()
    title.innerHTML = body
}
currentTrack()
window.setTimeout(currentTrack, 60 * 1000)
