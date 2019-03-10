<template>
  <v-container fluid>
    <v-toolbar flat color="white">
      <v-toolbar-title>Stretnutia</v-toolbar-title>
      <v-spacer />
    </v-toolbar>

    <v-layout wrap>
      <v-flex sm2 xs12 class="text-sm-left text-xs-center">
        <v-btn @click="$refs.calendar.prev()">
          <v-icon>keyboard_arrow_left</v-icon>
        </v-btn>
      </v-flex>
      <v-flex sm8 xs12 class="text-xs-center">
        {{ calendarTitle }}
      </v-flex>
      <v-flex sm2 xs12 class="text-sm-right text-xs-center">
        <v-btn @click="$refs.calendar.next()">
          <v-icon>keyboard_arrow_right</v-icon>
        </v-btn>
      </v-flex>
      <v-flex xs12 class="my-3">
        <v-sheet height="800">
          <v-calendar
            ref="calendar"
            type="month"
            locale="sk"
            :short-months="false"
            :weekdays="[1, 2, 3, 4, 5, 6, 0]"
            v-model="start"
            :end="end"
            class="elevation-1"
          >
            <template v-slot:day="{ date }">
              <template v-for="event in eventsMap[date]">
                <v-chip small :key="event.ID" @click="editSession(event)">
                  {{ event.UserName }}
                </v-chip>
              </template>
            </template>
          </v-calendar>
        </v-sheet>
      </v-flex>
      <v-flex xs12 class="text-sm-right">
        <v-btn color="primary" @click="addSession">
          <v-icon>add</v-icon>
        </v-btn>
      </v-flex>
    </v-layout>
    <v-dialog v-model="dialog" max-width="500px" @keydown.esc="close" persistent>
      <v-card>
        <v-form ref="personform" v-model="valid" lazy-validation>
          <v-card-title>
            <span class="headline">Stretnutie</span>
          </v-card-title>
          <v-card-text>
            <v-container grid-list-md>
              <v-form v-model="valid" ref="form" lazy-validation>
                <v-layout wrap>
                  <v-flex xs12>
                    <v-autocomplete 
                      v-model="form.user"
                      ref="userSelect"
                      :label="$vuetify.t('Užívateľ')"
                      :items="users"
                      hide-no-data
                      item-value="ID"
                      item-text="Name"
                      :rules="[rules.required]"
                    />
                  </v-flex>
                  <v-flex xs4>
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
                  <v-flex xs4>
                    <v-menu
                      ref="menuTime"
                      v-model="menuTime"
                      :close-on-content-click="false"
                      :nudge-right="40"
                      :return-value.sync="form.time"
                      lazy
                      transition="scale-transition"
                      offset-y
                      full-width
                      max-width="290px"
                      min-width="290px"
                    >
                      <template v-slot:activator="{ on }">
                        <v-text-field
                          v-model="form.time"
                          label="Čas"
                          prepend-icon="access_time"
                          readonly
                          v-on="on"
                        />
                      </template>
                      <v-time-picker
                        v-if="menuTime"
                        v-model="form.time"
                        format="24hr"
                        full-width
                        @click:minute="$refs.menuTime.save(form.time)"
                      />
                    </v-menu>
                  </v-flex>
                  <v-flex xs4>
                    <v-text-field
                      v-model="form.duration"
                      label="Trvanie (min)"
                      :rules="[rules.required]"
                      type="number"
                      mask="###"
                    /> 
                  </v-flex>
                  <v-flex xs12>
                    <v-textarea v-model="form.desc" :label="'Popis'" rows="4" auto-grow />
                  </v-flex>
                  <v-flex xs12>
                    <v-autocomplete 
                      v-model="form.persons"
                      multiple
                      :label="$vuetify.t('Účastníci')"
                      :items="persons"
                      hide-no-data
                      item-value="value"
                      item-text="text"
                      :rules="[rules.required]"
                    />
                  </v-flex>
                </v-layout>
              </v-form>
            </v-container>
          </v-card-text>

          <v-card-actions>
            <v-spacer />
            <v-btn color="blue darken-1" flat @click="close">
              Zrušiť
            </v-btn>
            <v-btn color="blue darken-1" flat @click="save">
              Uložiť
            </v-btn>
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
    dialog: false,
    rules: {
      required: value => !!value || 'Required.',
    },
    valid: true,
    menuDatePicker: false,
    menuTime: false,
    formDate: null,
    users: [],
    persons: [],
    events: [],
    form: {}
  }),
  computed: {
    calendarTitle: function() {
      return moment(this.start)
        .locale('sk')
        .format('MMMM YYYY')
    },
    // convert the list of events into a map of lists keyed by date
    eventsMap: function() {
      const map = {}
      // this.events.forEach(e => (map[e.DtoDatum] = map[e.DtoDatum] || []).push(e))
      this.events.forEach(function(e) {
        let ret = e.DtoDatum.substr(0, 10)
        return (map[ret] = map[ret] || []).push(e)
      })
      return map
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
      this.form.id = null
      this.form.user = this.$store.state.id
      this.formDate = moment().format('YYYY-MM-DD')
      this.form.dateFormatted = moment().format('DD.MM.YYYY')
      this.form.time = '09:00'
      this.form.desc = ''
      this.form.duration = 30
      this.form.persons = []
      this.dialog = true
    },
    editSession: function(event) {
      this.form.id = event.ID
      this.form.user = event.UserID
      this.formDate = event.DtoDatum.substr(0, 10)
      this.form.dateFormatted = moment(this.formDate, 'YYYY-MM-DD').format('DD.MM.YYYY')
      this.form.time = event.DtoDatum.substr(11, 5)
      this.form.desc = event.Description
      this.form.duration = event.Duration
      this.form.persons = event.Persons
      this.dialog = true
    },
    close: function() {
      this.dialog = false
    },
    save: function() {
      if (!this.$refs.form.validate()) {
        return
      }
      var params = {
        ID: this.form.id,
        UserId: this.form.user,
        Persons: this.form.persons,
        Duration: this.form.duration,
        Description: this.form.desc,
        DtoDatum: this.form.dateFormatted + ' ' + this.form.time
      }
      this.$axios.post("/r/api/session",params).then(response => {
        if (response.status == 202) {
          this.$store.commit("alert","Chyba: " + response.data.Error)
        } else { //OK
            // response.data
          this.dialog = false
          this.readSessions()
        }
      }).catch(response => {
        console.log('WRONG',response)
        this.$store.commit("alert","Chyba: " + response)
      })
    },
    readSessions: function() {
      this.$axios
        .get("/r/api/sessions")
        .then(response => {
          console.log("OK", response);
          if (response.status == 202) {
            this.$store.commit("alert", "Chyba: " + response.data.Error);
          } else {
            this.events = response.data;
          }
        })
        .catch(response => {
          console.log("WRONG", response);
          this.$store.commit("alert", "Chyba: " + response);
        });
    },
    readUsers: function() {
      this.$axios
        .get("/r/api/users")
        .then(response => {
          console.log("OK", response);
          if (response.status == 202) {
            this.$store.commit("alert", "Chyba: " + response.data.Error)
          } else {
            this.users = response.data
          }
        })
        .catch(response => {
          console.log("WRONG", response)
          this.$store.commit("alert", "Chyba: " + response)
        });
    },
    readPersons: function() {
      this.$axios
        .get("/r/api/persons")
        .then(response => {
          console.log("OK", response)
          if (response.status == 202) {
            this.$store.commit("alert", "Chyba: " + response.data.Error)
          } else {
            this.persons = response.data
          }
        })
        .catch(response => {
          console.log("WRONG", response)
          this.$store.commit("alert", "Chyba: " + response)
        });
    }
  },

  mounted: function() {
    console.log("SESSIONS.VUE, mounted")
    this.readUsers()
    this.readPersons()
    this.readSessions()
  }
};
</script>

<style scoped>
</style>
