document.addEventListener("DOMContentLoaded", () => {
    // Add event listener to the "Submit Vote" button
    document.getElementById("submitVote").addEventListener("click", submitVote);

    // Function to handle vote submission
    function submitVote() {
        const voterID = document.getElementById("voterID").value.trim();
        const candidate = document.getElementById("candidate").value.trim();

        // Validate input fields
        if (!voterID || !candidate) {
            alert("Both Voter ID and Candidate are required!");
            return;
        }

        const votePayload = {
            voterID: voterID,
            candidate: candidate,
        };

        fetch("http://localhost:8080/vote", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(votePayload),
        })
            .then((response) => {
                if (!response.ok) {
                    return response.text().then((text) => {
                        throw new Error(text || "Unknown error occurred");
                    });
                }
                return response.json();
            })
            .then((data) => {
                alert(data.message || "Vote successfully recorded!");
            })
            .catch((error) => {
                alert("An error occurred: " + error.message);
            });
    }
});
