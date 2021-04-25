export default {
  methods: {
    courseColorCode: (course) => {
      const map = {
        breakfast: 'is-warning',
        main: 'is-info',
        pudding: 'is-danger',
      };
      return map[course.name];
    },
  },
};
