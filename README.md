# Cheat Inspector client
### This project is submitted in [Ably hackathon](https://devpost.com/software/cheat-inspector)
## Inspiration

Cheat-Inspector is inspired by the need to maintain fairness and integrity in online hackathons and competitions. Plagiarism and project reuse have been significant issues that can dampen the spirit of such events.

Cheat-Inspector is a hackathon sentry that allows organizers to provide a fair competing platform in online events. Since plagiarism and re-use are the major problems in such submissions which ruins the spirit of hackathons, It utilizes intelligent algorithms to calculate project entropy and snapshots of participants' projects in real-time and visualizes the same for the organizers as a live graph in a pleasant user interface.

## What it does

Cheat-Inspector consists of multiple components:

- Cheat-Inspector-Server : web server written in TypeScript to act as the service to expose data for Cheat-Inspector-dashboards. Provides routes which utilize Redis-JSON to return data for business logic. Handle user creation, team creation, and reteriving the same from dashboard.
- Cheat-Inspector-Client : CLI written in Golang utilizing go-routines for performance ⚡ which calculates project entropy and snapshots and emits them to Ably
- Ably : Ably is the platform to power synchronized digital experiences in realtime. we are using ably to store, publish, synchronize and manage realtime data which is being emited from Cheat-Inspector-Clients
- Cheat-Inspector-Dashboard : a responsive and dynamic single-page application build using React and TailWind CSS, designed in a monochrome and minimal UI to focus on important data. Also provisions realtime graphs which render live feed of project status.

## Cheat-Inspector Client

- **The Binary**

  - Can be compiled for any platform, any architecture as far as GoLang supports it.
  - Generates unique signature for each device, which cannot be altered by changing configurations or be spoofed. So single machine cannot act as multiple devices.
  - Does **NOT** require admin privileges.
  - Device signatures are hardware independent, as they can be easily spoofed by VMs. MAC and BIOS settings are also ignored as they can be easily manipulated.
  - Automatically identifies the platform to display in the admin panel.

- **The Interface**

  - Interactive command line interface with option to navigate using arrow keys and validation check indicators builtin. If the validation is about to fail, the terminal shows red with an error message and there's no need to work-up again and again.
  - Secure TeamID Input : Since team IDs are used to join a team, the interface masks the input with `*` to add a layer of security in the user interface.

- **Configuration**

  - To make it effortless for users to use the application, server credentials can be embedded into the binary itself by the organizers (single point configuration declarations), and the participants can directly run the same.
  - In case if there are changes in server deployments, or the organizers come up with alternative servers to relay the updates, a configuration file can be used to declare the endpoints.
  - The configuration file should be named **cheatInspector.yaml** and be placed in the same directory as the binary sits.
  ```yml
    app:
      server: "https://some-fancy-server/api"
      debug: true

    ignore:
      - "public"
      - "cache"
      - "build"
      - "some-custom"
  ```
  - Restarting the client will first check if a configuration file is present. If its not, then display a message to ensure that the user knows, and go ahead to load the default configurations.
  - **Ignore Custom Directories** : since there can by multiple directories which contain auto-generated codebase, depending on the tech stack being used, therefore there is a provision to list all of them in the config file in the **ignore** section. Cheat-Inspector would then ignore those directories while calculating snapshots and entropy.

- **Functionality**

  - Allows user to create a team of people working together in an event / competition.
  - Allows user to join an existing team
  - Allows user to register their device
  - Thanks to unique device signatures, can identity old devices and avoid duplicate logins.
  - Walks through the project to calculate entropy and snapshot scores and feeds them to api servers.

- **Fancy Tech**
  - Written implementing go-routines for concurrency ⚡ and speed, utilizes minimal system resources, non blocking.
  - Recursively walks the the directory tree and creates a hashmap of the project in the memory.
  - Calculates the snapshot score and entropy of all the files in each iteration.
  - Snapshot score depends upon the file contents, length and size.
  - Entropy is a name given to diff-match score, which is basically the number of insertions and deletions required to move from one state to another.
  - The entropy is calculated using the [golang port](https://pkg.go.dev/github.com/sergi/go-diff/diffmatchpatch) of [Neil Fraser's google-diff-match-patch](https://github.com/google/diff-match-patch) code
  - Written in Golang, can be compiled to native binary for any operating system and architecture.

#### Vide demo: https://youtu.be/mGo2E-eCO94?si=kWaAj-HPKv-u0T6z
