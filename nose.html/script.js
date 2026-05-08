console.log("JavaScript cargado correctamente");

// Animación al hacer scroll
const sections = document.querySelectorAll("section");

const observer = new IntersectionObserver(entries => {
    entries.forEach(entry => {
        if (entry.isIntersecting) {
            entry.target.classList.add("show");
        }
    });
});

sections.forEach(section => {
    section.classList.add("hidden");
    observer.observe(section);
});

// Navbar dinámico
const nav = document.querySelector("nav");

window.addEventListener("scroll", () => {
    if (window.scrollY > 50) {
        nav.style.backgroundColor = "#d6cbb7";
    } else {
        nav.style.backgroundColor = "transparent";
    }
});

// Modo oscuro
const themeButton = document.getElementById("theme-toggle");

// Aplicar tema guardado
if (localStorage.getItem("theme") === "dark") {
    document.body.classList.add("dark-mode");
    themeButton.textContent = "Modo Claro";
}

themeButton.addEventListener("click", () => {
    document.body.classList.toggle("dark-mode");

    if (document.body.classList.contains("dark-mode")) {
        themeButton.textContent = "Modo Claro";
        localStorage.setItem("theme", "dark");
    } else {
        themeButton.textContent = "Modo Oscuro";
        localStorage.setItem("theme", "light");
    }
});