// web/static/js/main.js
document.addEventListener('DOMContentLoaded', () => {
    const searchInput = document.createElement('input');
    searchInput.type = 'text';
    searchInput.placeholder = 'Search artists...';
    searchInput.classList.add('search-input');

    const header = document.querySelector('header');
    header.appendChild(searchInput);

    const artistCards = document.querySelectorAll('.artist-card');

    searchInput.addEventListener('input', (e) => {
        const searchTerm = e.target.value.toLowerCase();

        artistCards.forEach(card => {
            const artistName = card.querySelector('h2').textContent.toLowerCase();
            if (artistName.includes(searchTerm)) {
                card.style.display = 'block';
            } else {
                card.style.display = 'none';
            }
        });
    });
});