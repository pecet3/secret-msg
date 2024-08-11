alert("hello")
const linkHolder = document.getElementById("linkHolder")
const msgForm = document.getElementById("msgForm")
console.log(msgForm.value(""))
msgForm.addEventListener('submit', async function (e) {
    e.preventDefault();

    fetch('/message', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ content: message })
    })
        .then(response => response.json())
        .then(data => {
            linkHolder.textContent = `${window.location.origin}/message/${data.token}`;
        })
        .catch((error) => {
            console.error('Error:', error);
            linkHolder.innerText()
        });
});