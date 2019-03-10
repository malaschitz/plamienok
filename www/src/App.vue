<template>
  <v-app id="keep">
    <v-navigation-drawer
      v-model="drawer"
      fixed
      clipped
      class="grey lighten-4"
      app
    >
      <v-list
        dense
        class="grey lighten-4"
      >
        <template v-for="(item, i) in items">
          <v-layout
            v-if="item.heading"
            :key="i"
            row
            align-center
          >
            <v-flex xs6>
              <v-subheader v-if="item.heading">
                {{ item.heading }}
              </v-subheader>
            </v-flex>
          </v-layout>

          <v-divider
            v-else-if="item.divider"
            :key="i"
            dark
            class="my-3"
          />
          <v-list-tile
            v-else
            :key="i"
            :to="item.link"
          >
            <v-list-tile-action>
              <v-icon>{{ item.icon }}</v-icon>
            </v-list-tile-action>
            <v-list-tile-content>
              <v-list-tile-title class="grey--text">
                {{ item.text }}
              </v-list-tile-title>
            </v-list-tile-content>
          </v-list-tile>
        </template>
      </v-list>
    </v-navigation-drawer>
    <v-toolbar
      color="amber"
      app
      absolute
      clipped-left
    >
      <v-toolbar-side-icon @click="drawer = !drawer" />

      <v-img
        :src="require('./assets/logo-plamienok.png')"
        contain
        max-height="55"
        max-width="140"
        @click="dashboard()"
      />

      <span class="title ml-3 mr-5">
        Plamienok&nbsp;<span class="font-weight-light">
          App
        </span>
      </span>
      <v-text-field
        solo-inverted
        flat
        hide-details
        label="Hladať"
        prepend-inner-icon="search"
      />
      <v-spacer />
      <v-menu
        v-if="$store.state.logged"
        :nudge-width="100"
      >
        <v-toolbar-title slot="activator">
          <span>{{ $store.state.name }} ({{ $store.state.email }})</span>
          <v-icon dark>
            arrow_drop_down
          </v-icon>
        </v-toolbar-title>

        <v-list>
          <v-list-tile @click="logout()">
            <v-list-tile-action>
              <v-icon>exit-to-app</v-icon>
            </v-list-tile-action>
            <v-list-tile-content>
              <v-list-tile-title>Odhlásiť</v-list-tile-title>
            </v-list-tile-content>
          </v-list-tile>
        </v-list>
      </v-menu>
    </v-toolbar>
    <v-content>
      <router-view />
    </v-content>
  </v-app>
</template>

<script>

    export default {
        name: 'App',

        components: {
        },

        data: () => ({
            drawer: null,
            items: [
                { heading: 'Home care' },
                { icon: 'child_care', text: 'Deti', link: '/persons'},
                { icon: 'done', text: 'Ciele', link: '/tasks'},
                { icon: 'lightbulb_outline', text: 'Návštevy' },
                { icon: 'touch_app', text: 'Telefón' },
                { divider: true },
                { heading: 'Poradňa' },
                { icon: 'calendar_today', text: 'Stretnutia', link: '/sessions' },
                { divider: true },
                { heading: 'Štatistiky' },
                { icon: 'settings', text: 'Hodiny' },
                { icon: 'chat_bubble', text: 'Autá' },
                { heading: 'Administrácia' },
                { icon: 'receipt', text: 'Lieky', link: '/meds' },
                { icon: 'pan_tool', text: 'Diagnózy', link: '/diagnoses' },
                { icon: 'group', text: 'Užívatelia', link: '/users' },
                { icon: 'directions_car', text: 'Autá', link: '/cars' },
                { icon: 'phonelink', text: 'Exporty' },
            ]
        }),

        methods: {
          dashboard: function() {
            return this.$router.push('/');
          },

          logout: function() {
            this.$store.commit("logout");
            this.$axios.defaults.headers.common["token"] = "";
            localStorage.clear();
            this.dashboard();
          },
        },

        computed: {
            alertMessage: function() {
              return this.$store.state.alert
            }
        },

        watch: {
            alertMessage: function(val) {
              if (val !== "") {
                console.log("ALERT", val);
              }
              this.$store.commit("alert","");
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
        },
    }
</script>

<style lang="stylus">
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

