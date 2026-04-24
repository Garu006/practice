//scroll suave para los enlaces del menú de navegación

document.querySelectorAll("nav a").forEach(link => {
    link.addEventListener("click", function(e) {
        e.preventDefault();

        const targetId = document.querySelector(this.getAttribute("href"));
        const targeetElement = document.querySelector(targetId);

        targetId.scrollIntoView({
            behavior: "smooth",
            block: "start",
            inline: "nearest",
        });
    });
});

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