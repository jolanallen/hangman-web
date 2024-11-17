function TestLetter(lettre) {
    alert('lettre taper est :' + lettre)
    const data = { message: lettre };

    fetch('http://localhost:8080/test', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
    .then(response => response.json())
    .then(data => {
        console.log('Réponse du serveur Go :', data);
    })
    .catch(error => {
        console.error('Erreur lors de l\'envoi :', error);
    });
}
