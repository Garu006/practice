//Animación al hacer scroll
const sectioon = document.querySelectorAll("Section");

const observer = new IntersectionObserver(entries => {
    entries.forEach(entry => {
        if (entry.isIntersecting) {
            entry.target.classList.add("visible");        
        }
    });
});

section.forEach(section => {
    observer.observe(section);
    section.classList.add("hidden");
});

// navbar dinamico
const nav = document.querySelector("nav");

window.addEventListener("Scroll", () => {
    if (window.scrollY > 50) {
        nav.style.backgroundColor = "#d6cbb7";
    } else {
        nav.style.backgroundColor = "transparent";
    }
});

// modo oscuro
const themeButton = document.getElementById("theme-toogle");

themeButton.addEventListener("click", () => {
    document.body.classList.toggle("dark-mode");
    
    if (document.body.classList.contains("dark-mode")) {
        themeButton.textContent = "Modo Claro";

    } else {
        themeButton.textContent = "Modo Oscuro";
    }
});