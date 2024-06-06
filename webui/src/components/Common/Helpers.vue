<script>
export default {
  methods: {
    getColor(item) {
      if (
        item === "Successful" ||
        item === "Success" ||
        item === "Unused" ||
        item === "Enabled"
      )
        return "success";
      if (item === "Reserved") return "warning";
      else return "error";
    },
  },
};
</script>