<template>
  <v-container fluid>
    <v-toolbar flat color="white">
      <v-toolbar-title>Stretnutia</v-toolbar-title>
      <v-spacer/>
    </v-toolbar>

    <v-layout wrap>
      <v-flex sm2 xs12 class="text-sm-left text-xs-center">
        <v-btn @click="$refs.calendar.prev()">
          <v-icon>keyboard_arrow_left</v-icon>
        </v-btn>
      </v-flex>
      <v-flex sm8 xs12 class="text-xs-center">{{ calendarTitle }}</v-flex>
      <v-flex sm2 xs12 class="text-sm-right text-xs-center">
        <v-btn @click="$refs.calendar.next()">
          <v-icon>keyboard_arrow_right</v-icon>
        </v-btn>
      </v-flex>
      <v-flex xs12 class="my-3">
        <v-sheet height="500">
          <v-calendar
            ref="calendar"
            type="month"
            locale="sk"
            :short-months="false"
            :weekdays="[1, 2, 3, 4, 5, 6, 0]"
            v-model="start"
            :end="end"
            class="elevation-1"
          ></v-calendar>
        </v-sheet>
      </v-flex>
      <v-flex xs12 class="text-sm-right">
        <v-btn color="primary" @click="addSession">
          <v-icon>add</v-icon>
        </v-btn>
      </v-flex>
    </v-layout>
    <v-dialog v-model="dialog" max-width="500px">
      <v-card>
        <v-form ref="personform" v-model="valid" lazy-validation>
          <v-card-title>
            <span class="headline">Stretnutie</span>
          </v-card-title>
          <v-card-text>
            <v-container grid-list-md>
              <v-layout wrap>
                <v-flex xs12 sm8 md6>
                  <v-text-field v-model="form.user" label="Užívateľ" :rules="[rules.required]" required />
                </v-flex>
                <v-flex xs12 sm6 md4>
                  <v-menu
                    v-model="menuDatePicker"
                    :close-on-content-click="false"
                    :nudge-right="40"
                    lazy
                    transition="scale-transition"
                    offset-y
                    full-width
                    min-width="290px"
                  >
                    <template v-slot:activator="{ on }">
                      <v-text-field
                        v-model="form.dateFormatted"
                        label="Dátum"
                        prepend-icon="event"
                        @change="formDate = parseDate(form.dateFormatted, formDate)"
                        v-on="on"
                        :rules="[rules.required]"
                        required
                      /> 
                    </template>
                    <v-date-picker 
                      v-model="formDate"
                      @input="menuDatePicker = false"
                      locale="sk"
                    />
                  </v-menu>
                </v-flex>
              </v-layout>
            </v-container>
          </v-card-text>

          <v-card-actions>
            <v-spacer/>
            <v-btn color="blue darken-1" flat @click="close">Zrušiť</v-btn>
            <v-btn color="blue darken-1" flat @click="save">Uložiť</v-btn>
          </v-card-actions>
        </v-form>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script>
import moment from "moment";

export default {
  name: "Sessions",

  data: () => ({
    start: moment().format("YYYY-MM-DD"),
    end: moment().month(11).date(31).format("YYYY-MM-DD"),
    items: [],
    dialog: false,
    rules: {
      required: value => !!value || 'Required.',
    },
    valid: true,
    menuDatePicker: false,
    formDate: null,
    form: {}
  }),
  computed: {
    calendarTitle: function() {
      return moment(this.start)
        .locale('sk')
        .format('MMMM YYYY')
    }
  },
  watch: {
    'formDate': function(newValue){
      console.log('formDate watcher', newValue)
      this.form.dateFormatted = moment(newValue).format('DD.MM.YYYY')
    }
  },
  methods: {
    parseDate: function(aFormattedDatum, aOldValue) {
      console.log('parseDate', aFormattedDatum, aOldValue)
      let ret = aOldValue
      if (aFormattedDatum !== undefined && moment(aFormattedDatum,'DD.MM.YYYY').isValid()) {
        ret = moment(aFormattedDatum,'DD.MM.YYYY').format('YYYY-MM-DD')
      }
      return ret
    },
    addSession: function() {
      this.form.user = this.$store.state.name
      this.formDate = moment().format('YYYY-MM-DD')
      this.form.dateFormatted = moment().format('DD.MM.YYYY')
      this.dialog = true
    },
    close: function() {
      this.dialog = false
    },
    save: function() {
      this.dialog = false
    },
    readData: function() {
      this.$axios
        .get("/api/sessions")
        .then(response => {
          console.log("OK", response);
          if (response.status == 202) {
            this.$store.commit("alert", "Chyba: " + response.data.Error);
          } else {
            this.items = response.data;
          }
        })
        .catch(response => {
          console.log("WRONG", response);
          this.$store.commit("alert", "Chyba: " + response);
        });
    }
  },

  mounted: function() {
    console.log("SESSIONS.VUE, mounted");
    // this.readData();
  }
};
</script>

<style scoped>
</style>
