## Deep diving into Interfaces

## Implementation idea

The Scenario: The Weather Alert System
You are building a weather monitoring application. When the system detects extreme weather (like a tornado or blizzard), it needs to alert the users.
Right now, the product manager only wants to send alerts via Email, but they have already warned you that next sprint, they will want to add SMS text alerts.
Because you are a forward-thinking engineer, you want to build the WeatherStation (the executor) so that it is completely decoupled from the specific alerting mechanism.
