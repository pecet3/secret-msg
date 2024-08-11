const linkHolder = document.getElementById("linkHolder")
const msgForm = document.getElementById("msgForm")

msgForm.addEventListener('submit', async function (e) {
    e.preventDefault();

    const formData = new FormData(msgForm)
    const message = formData.get("message")
    const response = await fetch('/message', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ message: message })
    })
    if (!response.ok) {
        alert("Something went wrong... Try again")
        return
    }
    const result = await response.json()
    alert(result.token)
});