document.getElementById('voteForm').addEventListener('submit', function(e) {
    e.preventDefault();
    const candidate = document.getElementById('candidate').value;

    fetch('http://localhost:8080/vote', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ candidate })
    })
    .then(response => response.json())
    .then(data => {
        document.getElementById('status').innerText = 'Vote submitted successfully!';
    })
    .catch(err => {
        document.getElementById('status').innerText = 'Error submitting vote';
    });
});
