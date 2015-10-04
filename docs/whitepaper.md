# Overall System Design

### Version `2015.10.03b-alpha`

This is a whitepaper which details the system and user UI flow design of Notion, with a focus on the backend. This document formalizes and version controls the discussion held on the night of 30-September. It is currently open for feedback, with the first non-alpha version slated for finalization early next week.

# Version Log

* `2015.10.03b-alpha`: Expanded information about what can be sent over websockets.
* `2015.10.03a-alpha`: Initial publish. Signed off: mhoc

# Concepts

These are fundamental concepts in the notion system. The list of fields on these shouldn't be considered final or complete; that list can be found in the `src/model` folder.

* User: A single user of the notion system. Users are assigned a unique Id by Notion, hence referred to as `user_id`. Users are currently authenticated through Facebook, but in the future might be authenticated through additional means. Each user has a role, which is currently always `student` but might in the future be `instructor` or `administrator` depending on future features. Each user can be 'verified' or not, which currently has no meaning but might in the future.

* School: A university (or other learning institution) where classes might be associated with. Metadata about schools is very stark right now, besides having an id, name, and location. The location is currently a simple user-friendly string, but might be expanded later.

* Course: A 'class designation', but not a specific instance of a class. For example: CS 490 Software Engineering Senior Project. Courses have an id, number (CS 490), and name (Senior Project). A course has a N:1 relationship with a school.

* Section: An instance of a course. Sections have a notion-generated unique id, a professor, a year, a semester (Fall, Spring, etc) and some other form of unique course identifier which we are referring to as a CRN.

* Notebook: A collection of topics, each of which contains notes. A notebook has a privacy setting which may be set to 'private' (only the owner can see it) or 'section' (public, associated with a single section). If a notebook is private, it has an N:1 relationship with a user, designated as the 'owner'. Every section has a single notebook associated with it. This section notebook has no owner (until the concept of 'instructors' is implemented, if ever) and can be considered public (until the concept of allowing instructors to verify students is implemented, if ever). Users can 'subscribe' to notebooks, which indicates that the notebook will appear in their list of notebooks. Users are automatically subscribed to any notebook they own.

* Topic: A collection of notes centered around a common theme. This theme might be a lecture on a given day, a chapter in the textbook, or anything else the users want. Topics are student-created and named (until the concept of instructors verifying/locking topics is implemented, if ever). The recommendation engine operates at the scale of topics; any note which exists in topic N will receive recommendations from other users' notes in topic N. Topics have an id, name, and a date-created field for the UI to sort on.

* Note: A markdown-formatted text document. A note always has a single owner, including in 'section' notebooks. Notes may only be edited by their owner (until collaborative editing is implemented, if ever). Notes in 'section' notebooks are readable by the public. Notes in 'private' notebooks are not. Notes must always belong to a topic.

* Changeset: A single atomic change to a note. A note could be reverse-engineered by traversing an ordered list of changesets. The exact format of a changeset is still under-development.

# Combining Sections

In some courses it might make sense to combine sections. For example, if there are two sections of CS180 taught by the same professor, there is value in combining them. However, in most cases it does not make sense to automatically combine sections.

Considering this, for the time being combining sections will not be implemented in any capacity. Users wishing to replicate this functionality will simply have to 'subscribe' to both sections' notebooks.

# User Creation

Account creation and login is both done through the `login` endpoint. This endpoint requires the inclusion of an authMethod parameter, which currently can only be 'facebook'.

In the case of facebook, when this endpoint is hit the API first validates the access_token with the facebook graph API, which returns the user's name and user id. The API also makes a call to convert the access_token to a long-lived token. The API then checks postgres for a user with the given facebook user id. If they exist, a `202 ACCEPTED` is returned to the client with the user's notion ID and a notion auth token. If they don't exist, they are created and a `201 CREATED` is returned with the same information.

# User Authentication

All authentication is handled through an authentication token returned from the `login` endpoint. This token must be included with every API call (except `login`), or else a `401` will be returned. This token never passively expires, but can be invalided with a call to the `logout` endpoint.

# Facebook Auth Token

Facebook auth tokens do expire. The API will convert any API token it receives on `login` to a long-lived token, which should last at least a month. However, this is only because the default TTL is around 2 hours, and its possible a user could use our app for that long without refreshing the page. Clients should always re-contact facebook for a new auth token and make a new call to `login` to receive a new notion token every time a page is refreshed or the mobile app is loaded to ensure active users always have a valid facebook token.

That being said, the downside to having an invalid facebook token is currently almost nil because all contact with the graph API is done on `login`. This may change in the future, however.

# REST API Calls

Restful API calls are provided for all models to receive a 'snapshot' of data at the time of the call. In many cases, the data these calls are accessing is very static and thus safe (list of sections, schools, etc). However, as one example, a call to `GET /note` might not always return the most up-to-date note information due to the fact that note content on notes currently being edited is cached in redis and only flushed at frequent intervals.

These calls are all authenticated, and the content returned from them will only include resources which the current user is authenticated to see.

# Websocket API (UNDER DEVELOPMENT)

_Development Note: The term 'websocket' is used consistently here, but this could refer to a yet-undecided websocket implementation including normal websockets, socket.io, or browsersocket._

Real-time functions of this system are better accessed through the websocket API. Websockets are bound on a per-notebook basis; when a user opens a notebook, the client should request a websocket connection for that notebook through a REST API call. If the notebook is 'private' and the user is not authenticated to see it, this call will fail.

Several different things can be sent over a websocket. At all times, clients may receive information about new topics and new notes being added to a notebook by other users. When a note is open, clients can expect to receive information about note recommendations from other notes in the same topic. In the future, clients may also receive changesets on the note they are editing if collaborative editors are present.

# Creating a Topic

Any user which has access to a notebook has the ability to create topics. This can only be done through a REST api call. You will be given back an id for that topic.

# Creating a Note

Notes must be created 'underneath' a topic and must be done through a REST API call by essentially requesting that a new empty note be created for this user. You will be given back an id for this note. Depending on the client implementation, you may update the note through simple REST API calls, but the websocket API is recommended.

# Opening a Note

In well-designed client implementations, a websocket connection will already be open when the request to create a note is processed. However, you must send a websocket request to 'open' a note. This will subscribe your websocket connection to receive recommendations for the topic the note belongs to and changesets for that note from other users (in future implementations).

# Closing Notes

When the user navigates away from the note, clients must send a 'close note' request over the websocket which will de-subscribe your websocket from recommendation and changeset updates.

# Closing Notebooks

When the user navigates away from an entire notebook, clients must close the websocket entirely.

# Changesets

A changeset is defined as a single atomic change to a note. If updating notes through the REST endpoint, the concept of changesets is ignored and instead you simply PUT the entire content of the note all at once. Through a websocket, clients must send changesets which the server then processes and uses to update the master copy of their note. The interval at which changesets are sent is up to the client, but it should be on the order of every update to the note. Clients may attempt some form of caching such that changesets are not sent on every note update, but be aware that this might hurt UX when collaborative editing is added.

The exact content of a changeset is under development.

Clients should anticipate the addition of a feature in which they will receive changesets for an open note from 'collaborative' users. Clients must be able to integrate these changesets into their view of the note.
