# Ideas

Just some ideas and possible future features.

* 'goatd server' should validate and fix files only during startup and later
  upon explicit HTTP requests. Otherwise we increase the risk of race
  conditions because the user could be executing CLI commands at the same time.
  Since the tool is single user in general we don't need to worry about race
  conditions. It's up to the user to ensure they are not overwriting their own
  changes. The GUI must reload information in regular intervals to be up to
  date.
* Do projects need an ID? Probably not. If they do, the MD file could contain a
  h1 Meta section at the end of the file where such information is stored.
* Projects should have a short name (file name of the MD file) and a long name
  which is the title of the MD document. Unless we use IDs the short name
  should be used e.g. to add a task. Limitations upon the short name should
  only be imposed by the file system (e.g. spaces are fine).
* Is it better to have 'goatd add X' where X is task, project etc. or 'goatd
  project add' etc.? I think the second feels more natural. 'hugo' does
  verb-noun and it corresponds to the way we talk in English.
* Projects could have subprojects. At the same time it may be nice to split
  projects into multiple files and subdirectories, in particular if the MD
  directory should be used not only for storing tasks but also more prose about
  the project. Maybe the project could start as clean-room.md and if necessary
  be expanded into clean-room/index.md and later if appropiate into
  clean-room/\_SUBPROJECTS/new-shelf.md
* Tasks could be represented as "[] This is a task DUE:2022-05-23 START:2022-05-21 PRIO:60 TAGS:asdf,ghi"
* Task dependencies can be represented by nesting them in lists like Zim does
* Tasks should alway be at the beginning of the line or as part of a list?
* Do Tasks need an ID? It would make things more difficult. goatd could
  automatically add IDs to tasks that don't have them and verify there are no
  duplicates. there could be a global counter with the last assigned ID. But
  should only be done if really necessary.
* Tasks could have four states: [ ] open, [✓] U+2713 or [x] done, [-]
  won't do, [→] U+2192 or [>] moved somewhere else
  * GitHub is using x to mark the task as done:
    https://github.blog/2014-04-28-task-lists-in-all-markdown-documents/
  * The fancy Unicode icons don't seem to add any value. People using a text
    editor won't benefit in the GUI can translate transparently anyway.
  * Maybe it would be best to start simple with binary todo/done [ ] and [x]
    like GitHub, also for compatibility. The 'wont do' state is quite useful to
    me to document that I didn't do something and why, but maybe it becomes too
    complicated...
* How to represent 'waiting for'? A task state [o] or [🕐]? Or a tag? Something else?
  * Maybe a TAG is also a good option like Zim does
* Tasks can be defined anywhere in the project doc but when they are moved to
  another document they are put at the end of the h1 Tasks section at the end
  of the document
* Add a Reference/ folder for general reference documentation about anything
  and everything. i.e. convert this project into a general Wiki
* Search in the GUI
* CLI with ncurses or similar
* It should be possible to mark tasks for "today" e.g. with TAGS:today
* Tasks can have priority 0-100. Higher means more urgent. Default 40. Priority
  0 means it won't be listed anywhere except on the project page. This allows
  defining tasks for the future without swamping the overview.
* When moving pages e.g. from Projects to Done links should be updated
* Possible structure
Projects/
Done/
Trash/
SomeDay/
Inbox.md
Reference/
* When moving to Done, Trash or Reference a timestamp should be appended if
  there is a naming conflict.
* Tasks in an overview are only listed for 'Projects'
* It's possible to see tasks in the Inbox and move them to projects
* 'goatd c some free text' or 'goatd capture some free text' to add something
  to the inbox. Make capturing painless.
* goatd task add 'some text' --project asdf
* Priority of old tasks increases automatically every week? Age could be determined via 'git log'.
* Allow synchronizing automatically with a Git remote. Display a warning in the
  CLI and GUI if there is a merge conflict that needs manual resolution.
* Power user: Allow seeing some Git info via CLI or GUI (e.g. to check remote is reachable)
* Validation of the format is a must. Automatic changes seems like a good idea
  e.g. format everything, add missing sections and params to tasks (no don't,
  makes it more painful to edit the tasks manually)
* Use mage https://github.com/magefile/mage instead of Makefile
* Have an 'init' command that initializes a directory structure or have that
  happen automatically when specifying a new base-dir?
* Have a verify/validate command that checks and maybe fixes structural problems etc.


# Prior art

Some other projects that may serve as influence or turn out to be better than goatd:

* https://github.com/rockiger/akiee
* https://github.com/benjaminoakes/markdo
* https://www.omnigroup.com/omnifocus/features/
* https://www.taskpaper.com/
