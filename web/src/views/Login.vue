<template>
  <section class="hero is-fullheight">
    <div class="hero-body">
      <div class="container">
        <div class="columns is-centered">
          <div class="column is-5-tablet is-5-desktop is-4-widescreen">
            <AuthForm
              :has-google-auth="googleAuthEnabled"
              :is-expanded="isExpanded"
              :email="email"
              :password="password"
            />
            <b-notification
              position="is-top-right"
              v-if="showDemo"
              type="is-info is-light"
              aria-close-label="Don't show again"
              role="tooltip"
              @close="onClose"
            >
              First time here?
              Try <a @click.prevent="fillDemoData">demo</a> and see how it works!
            </b-notification>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import AuthForm from '@/components/AuthForm.vue';

export default {
  name: 'Login',
  components: { AuthForm },
  data() {
    return {
      email: '',
      password: '',
      showDemo: !localStorage.getItem('skip_demo'),
      isExpanded: false,
    };
  },
  computed: {
    googleAuthEnabled() {
      return !!window.config.MP_CLIENT_ID;
    },
  },
  methods: {
    fillDemoData() {
      this.isExpanded = true;
      this.email = 'demo@demo.com';
      this.password = 'demo';
    },
    onClose() {
      localStorage.setItem('skip_demo', '1');
    },
  },
};
</script>

<style scoped>

</style>
