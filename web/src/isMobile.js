import { ref } from 'vue';

const isMobile = ref(window.innerWidth <= 768);

window.addEventListener('resize', () => {
  isMobile.value = window.innerWidth <= 768;
}, { passive: true });

export default isMobile;
