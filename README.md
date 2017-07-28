# lowest-common-ancestor

## Assignment 

Bureaucr.at is a typical hierarchical organization. 
Claire, its CEO, has a hierarchy of employees reporting 
to her and each employee can have a list of other employees reporting to him/her. 
An employee with at least one report is called a Manager.

Your task is to implement a corporate directory for Bureaucr.at 
with an interface to find the closest common Manager (i.e. farthest from the CEO) 
between two employees. You may assume that all employees eventually report up to the CEO.

Here are some guidelines:
- Resolve ambiguity with assumptions.
- The directory should be an in-memory structure.
- A Manager should link to Employees and not the other way around.
- We prefer that you to use Go, but accept other languages too.
- How the program takes its input and produces its output is up to you.

## Algorithm 

1. The task might be represented by Lowest Common Ancestor problem.
In common case there are few different algorithms to solve this problems.

2. In this case Tarjan's off-line lowest common ancestors algorithm was chosen 
to solve this task. This algorithm includes [preprocessing](http://dl.acm.org/citation.cfm?id=321884) for O(α(n)) time complexity 
and gives the opportunity to provide constant-time queries after this preprocessing.

3. In the LCA problem a node might be an ancestor of itself. In our case it might be strange
to assume that Employees might report to themselves. 
So, if A is the Manager of B and B is the Manager of C assume that the closest common Manager
for B and C will be A, but not B.

4. However, to be able to find the Manager between CEO and any other Employee, 
assume that CEO might report herself. 
So, the Manager between the CEO and any other Employee will be the CEO.

5. Tarjan's algorithm represents a depth-first search with making of disjoint-sets and
coloring of traversed nodes.
To be able to provide (3) it's not enough to just color nodes, we also need to store information
about node immediate ancestors. Luckily, we can use the same variable to mark node as colored 
and to store its immediate ancestor. So, if the immediate ancestor of the node was stored,
assume that the node was colored.

## Implementation

1. The task is implemented as a web-service.

2. The organization's structure must be given as a JSON object represented by this JSON schema 
(see an example in `cmd/default_config.json`):

    
    {
        "title": "A directory",
        "type": "object",
        "properties": {
            "name": {
                "title": "Unique name of the current directory Manager",
                "type": "string"
            },
            "employees": {
                "title": "List of the Employees",
                "type": "array",
                "items": {
                    "$ref": "#"
                }
            },
        },
        "required": ["name"]
    }
    
3. To deal with nodes of the tree we need to use unique names.
Assume that names of all Employees are unique. Otherwise,
we need to use unique IDs or something like this.

4. Configurable ENV parameters of the service are:
- `PORT` to access a service via the given port
- `CONFIG_FILE` which gives a path to configuration file (2) to set a structure of Bureaucr.at company

5. There are no too many validation levels in this service. For example, there is no validation of
unique names in configuration file and there is no too much validation in API requests processing.
In production-ready case validation must be realized better.

6. Then the service starts, it makes a "matrix" of the names of the closest Managers.
Then the matrix is built, the service is ready to listen to the requests.
The matrix stores as in-memory structure represented by map.
The matrix is symmetric, so for each couple of the Employees there is only one entry in the matrix.
The matrix stores N * [N - 1] / 2 elements where N is a total count of the Employyes.

## Gettitng started 

We need [Go](https://golang.org) and [Glide](https://glide.sh) to be able to prepare the service.

    make build
    ENV PORT=8888 CONFIG_FILE=./cmd/default_config.json make run 

Request example:

    curl -i http://127.0.0.1:8888/api/v1/lowest-common-manager?employees=Faith,Ivo

## Demo

