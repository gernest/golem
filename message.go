/*

   Copyright 2013 Niklas Voss

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.

*/

package golem

// Message is container for unprepared data and therefore holds the event name and the pointer to the struct holding the data.
type Message struct {
	event string
	data  interface{}
}

func (m *Message) GetEvent() string {
	return m.event
}

func (m *Message) GetData() interface{} {
	return m.data
}

func (m *Message) SetData(d interface{}) {
	m.data = d
}

func (m *Message) SetEvent(evt string) {
	m.event = evt
}
