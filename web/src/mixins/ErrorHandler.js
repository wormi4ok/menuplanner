import { mapGetters } from 'vuex';

export default {
  computed: {
    ...mapGetters({
      errorMessage: 'getError',
    }),
  },
  watch: {
    errorMessage(data) {
      let message = '';
      if (Object.prototype.hasOwnProperty.call(data, 'errors')) {
        data.errors.forEach((error) => {
          message += `<p>${error.message}: ${error.field}</p>`;
        });
      } else {
        message = data;
      }
      this.$buefy.dialog.alert({
        title: 'Error',
        message,
        type: 'is-danger',
        hasIcon: true,
        icon: 'times-circle',
        iconPack: 'fas',
        ariaRole: 'alertdialog',
        ariaModal: true,
      });
    },
  },
};
