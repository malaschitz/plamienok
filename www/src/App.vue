<template>
  <v-app id="keep">
    <v-app-bar color="amber" app clipped-left>
      <v-app-bar-nav-icon @click="drawer = !drawer" />
      <v-img
        :src="require('./assets/logo-plamienok.png')"
        contain
        max-height="55"
        max-width="140"
        @click="dashboard()"
      />

      <span class="title ml-3 mr-5">
        Plamienok&nbsp;
        <span class="font-weight-light">App</span>
      </span>
      <v-text-field
        solo-inverted
        flat
        hide-details
        label="Hladať"
        prepend-inner-icon="mdi-magnify"
      />
      <v-spacer />

      <v-menu v-if="$store.state.logged" origin="center center" transition="scale-transition">
        <template v-slot:activator="{ on }">
          <v-btn text v-on="on">{{ $store.state.name }} ({{ $store.state.email }})</v-btn>
        </template>

        <v-list>
          <v-list-item @click="logout()">
            <v-list-item-title>Odhlásiť</v-list-item-title>
          </v-list-item>
        </v-list>
      </v-menu>
    </v-app-bar>

    <v-navigation-drawer v-model="drawer" app clipped color="grey lighten-4">
      <v-list dense class="grey lighten-4">
        <template v-for="(item, i) in items">
          <v-row v-if="item.heading" :key="i" align-center>
            <v-col cols="6">
              <v-subheader v-if="item.heading">{{ item.heading }}</v-subheader>
            </v-col>
          </v-row>

          <v-divider v-else-if="item.divider" :key="i" dark class="my-1" />

          <v-list-item v-else :key="i" :to="item.link">
            <v-list-item-action>
              <v-icon>{{ item.icon }}</v-icon>
            </v-list-item-action>
            <v-list-item-content>
              <v-list-item-title class="grey--text">{{ item.text }}</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </template>
      </v-list>
    </v-navigation-drawer>

    <v-content>
      <router-view />
    </v-content>

    <v-snackbar v-model="snackbar" :timeout="5000" :color="snackbarColor">
      {{ alertText }}
      <v-btn flat @click="snackbar = false">Zavrieť</v-btn>
    </v-snackbar>
  </v-app>
</template>

<script>
export default {
  name: "App",

  components: {},

  data: () => ({
    drawer: null,
    items: [
      { heading: "Home care" },
      { icon: "mdi-human-male-boy", text: "Deti", link: "/persons" },
      { icon: "mdi-target", text: "Ciele", link: "/tasks" },
      { icon: "mdi-cake", text: "Meniny", link: "/meniny" },
      { icon: "mdi-hospital-box", text: "Návštevy", link: "/visithomes" },
      { icon: "mdi-phone-classic", text: "Telefón", link: "/visitcalls" },
      { divider: true },
      { heading: "Poradňa" },
      { icon: "mdi-account-supervisor", text: "Stretnutia", link: "/sessions" },
      { divider: true },
      { heading: "Štatistiky" },
      { icon: "mdi-account-clock", text: "Hodiny" },
      { icon: "mdi-car-side", text: "Autá" },
      { heading: "Administrácia" },
      { icon: "mdi-pill", text: "Lieky", link: "/meds" },
      { icon: "mdi-stethoscope", text: "Diagnózy", link: "/diagnoses" },
      { icon: "mdi-account-group", text: "Užívatelia", link: "/users" },
      { icon: "mdi-car-multiple", text: "Autá", link: "/cars" },
      { icon: "mdi-file-table-box-multiple", text: "Exporty" }
    ],
    alertText: "",
    snackbar: false,
    snackbarColor: "info"
  }),

  methods: {
    dashboard: function() {
      return this.$router.push("/");
    },

    logout: function() {
      this.$store.commit("logout");
      this.$axios.defaults.headers.common["token"] = "";
      localStorage.clear();
      this.dashboard();
    }
  },

  computed: {
    alertMessage: function() {
      return this.$store.state.alert;
    },

    normalMessage: function() {
      return this.$store.state.message;
    }
  },

  watch: {
    alertMessage: function(val) {
      if (val !== "") {
        this.snackbar = true;
        this.alertText = val;
        this.snackbarColor = "error";
        console.log("ALERT:" + val);
      }
      this.$store.commit("alert", "");
    },

    normalMessage: function(val) {
      if (val !== "") {
        this.snackbar = true;
        this.alertText = val;
        this.snackbarColor = "success";
        console.log("MESSAGE:" + val);
      }
      this.$store.commit("message", "");
    }
  },

  mounted: function() {
    console.log("APP MOUNTED");
  },

  props: {
    source: {
      type: String,
      default: ""
    }
  }
};
</script>

<style>
#keep main .container {
  height: 660px;
}

.navigation-drawer__border {
  display: none;
}

.text {
  font-weight: 400;
}
</style>
