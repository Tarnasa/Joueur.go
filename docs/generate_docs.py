import json
import os
import os.path
from shutil import copyfile, rmtree
import subprocess
import sys
from time import sleep

OUTPUT_DIR = './output'
GAMES_DOCS_DIR = './games'

def run(command, ignore_errors=False, **kwargs):
    error_code = subprocess.call([command], shell=True, **kwargs)
    if not ignore_errors and error_code != 0: # an error happened
        print('unexpected error')
        sys.exit(error_code)
    return error_code

def ensure_clean_dir(dirs):
    if os.path.isdir(dirs):
        rmtree(dirs)
    os.makedirs(dirs)

def github_link_to(game_name):
    return 'https://github.com/siggame/Cadre/blob/master/Games/{}/'.format(game_name)

def package_link_to(game_name):
    return 'pkg/joueur/games/{}/index.html'.format(game_name.lower())

def replace_pkg_links(s):
    return s.replace("pkg.1", "")

def template_collapsible_section(header, contents):
    return """
<div id="{id}" class="toggleVisible">
	<div class="collapsed">
		<h2 class="toggleButton" title="Click to show {header} section">{header} ▹</h2>
	</div>
	<div class="expanded">
		<h2 class="toggleButton" title="Click to hide {header} section">{header} ▾</h2>
		<div class="{id}-contents">
			{contents}
		</div> <!-- .{id}-contents -->
	</div> <!-- .expanded -->
</div>""".format(
    id=header.lower().replace(' ', '-'),
    header=header,
    contents=contents,
)

def update_file(output_path, formatter):
    with open(os.path.join(OUTPUT_DIR, output_path), 'r+') as output_file:
        contents = output_file.read()
        contents = replace_pkg_links(contents)

        ih = contents.index('</head>')
        contents = '{}<link rel="icon" type="image/x-icon" href="{}favicon.ico" />{}'.format(
            contents[:ih],
            '../' * output_path.count('/'),
            contents[ih:],
        )

        if formatter:
            contents = formatter(contents)

        output_file.seek(0)
        output_file.write(''.join(contents))

run('go get golang.org/x/tools/cmd/godoc@v0.0.0-20191213221258-04c2e8eff935')

gopath = os.getenv('GOPATH')
godoc_process = subprocess.Popen(gopath + '/bin/godoc', cwd='../')

# wait 3 seconds for the above process to be ready,
# no easy API to ensure it actually is
sleep(3)

ensure_clean_dir(OUTPUT_DIR)
print('-> Going to scrape the godoc server, this will take some time...')
wget_error_code = run(
    'wget -m -k -q -erobots=off -X src/ --no-host-directories --no-use-server-timestamps http://localhost:6060',
    cwd=OUTPUT_DIR,
    timeout=300, # 5 min
    ignore_errors=True,
)

print('-> Done scraping. Killing process')
godoc_process.kill()

if wget_error_code not in [0, 8]: # 0 is ok, 8 is server error we don't care about
    print('!!-> wget error code', wget_error_code)
    # sys.exit(wget_error_code)

print('-> Injecting additional documentation into scraped html files')
games = {}
for filename in os.listdir(GAMES_DOCS_DIR):
    with open(os.path.join(GAMES_DOCS_DIR, filename), 'r') as game_data_file:
        parsed_file = json.load(game_data_file)
        games[parsed_file['game_name']] = parsed_file

# Inject/change up the index.html file a bit to be more Cadre game centric.
def root_index_update_contents(contents):
    # auto collapse all sections because the vast majority of the packages
    # are redundant to go users
    contents = contents.replace('class="toggleVisible"', 'class="toggle"')

    # and slice in some additional documentation data about the games
    i = contents.index('<h1>')
    return contents[:i] + """
<h1>Joueur.go Documentation</h1>

This is the documentation for the Go Cadre client and its various game
packages.

{}
{}
""".format(
    template_collapsible_section(
        'Games', """
<p>These are the games that are available to play via the Go Client. Their
 source code is stored in the directory: <code>games/game_name/</code>, where
 <code>game_name</code> is the name of the game.
</p>

<dl>
{}
</dl>""".format('\n'.join([
    """
        <dt><strong><a href="{pkg_link}">{game_name}</a></strong></dt>
        <dd>{description}</dd>
    """.format(
            game_name=game_name,
            pkg_link=package_link_to(game_name),
            description=games[game_name]['description']
        ) for game_name in sorted(games.keys())
        ]))
    ),
    template_collapsible_section(
        "Coding Your AI", """
<h3>Interfaces</h3>
<p>With the exception of your AI being a <code>struct</code>, all of the game
 components you will interact with are done through Go
 <code>interfaces</code>. This means that all attributes must be accessed via
 function calls, e.g:
<pre>
player_name := ai.Player().Name()
</pre>
</p>
<br/>
<p>Unless otherwise noted in the documentation, assume all interfaces are
populated by an instance of a struct implementing that interface. However some
attributes, function call, etc will explicitly tell you if the returned value
can be nullable (<code>nil</code> pointer).
</p>

<h3>Modifying non AI files</h3>
<p>
Each interface type inside of <code>games/game_name/</code>, except for your
 <code>ai.go</code> should ideally not be modified.<br />
They are intended to be read only constructs that hold the state of that
 object at the point in time you are reading its properties.
</p>
<p>
With that being is said, if you really wish to add functionality, such as
 helper functions, ensure they do not directly modify game state information,
 or interfere with our existing functionality, or there is a good chance your
 client will crash during gameplay with a <code>DELTA_MERGE_FAILURE</code>.
<p>
<p>Implimentation logic for the interfaces (except your AI) is all tucked away
 in <code>games/internal/game_name_impl</code>. It is highly reccomended not to
 modify these files as they are largey written by our <a href="{creer_link}">
 Creer code generation</a> tool and may need to be modified if the game
 structure is tweaked.
</p>

<h3>Game Logic</h3>

<p>If you are attempting to figure out how the logic is executed for a game,
 that code is <strong>not</strong> here.<br />
 All <a href="{cadre_link}">Cadre</a> game clients are dumb state tracking
 programs that facilitate IO between a game server and your AI in whatever
 programming language you choose.
</p>
<p>
If you wish to get the actual code for a game check in the
 <a href="{cerveau_link}">Cerveau game server</a>. Its directory structure is
 similar to most clients (such as this one).
<p>
""".format(
        cadre_link='https://github.com/siggame/Cadre',
        cerveau_link='https://github.com/siggame/Cerveau',
        creer_link='https://github.com/siggame/Creer',
    )),
) + contents[i:]

update_file('index.html', root_index_update_contents)

# for each game, add additional text explaining the game
for game_name, game_docs in games.items():
    def game_index_update_contents(contents):
        i = contents.index('<div id="pkg-index"')
        return contents[:i] + template_collapsible_section(
            game_name + " Game Info",
            """
<p>
{description}
</p>
<h3>More Info</h3>
<p>
The full game rules for {game_name} can be found on <a href="{github}rules.md">GitHub</a>.
</p>
<p>
Additional materials, such as the <a href="{github}story.md">story</a> and <a href="{github}creer.yaml">game template</a> can be found on <a href="{github}">GitHub</a> as well.
</p>
""".format(
            description=game_docs['description'],
            game_name=game_name,
            github=github_link_to(game_name)
        )
    ) + contents[i:]

    update_file(package_link_to(game_name), game_index_update_contents)

copyfile('./favicon.ico', os.path.join(OUTPUT_DIR, 'favicon.ico'))

print('<- Done generating Go docs')
