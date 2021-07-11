# Create security groups

A security group acts as a virtual _firewall_ that controls the traffic for one or more instances.

When you launch an instance, you associate one or more security groups with the instance.

You add rules to each security group that allow traffic to or from its associated instances.

You can modify the rules for a security group at any time; the new rules are automatically applied to all instances that are associated with the security group. When we decide whether to allow traffic to reach an instance, we evaluate all the rules from all the security groups that are associated with the instance.

## ELB (load balancer) security group

1. add these rules
   - HTTP TCP 80 Custom IP ```<load-balancer-sg Group ID>```
   - SSH TCP 22 My IP
2. copy __GROUP ID__
3. add this rule
   - MySQL TCP 3306 Custom IP ```<web-servers-sg Group ID>```

---

# Load Balancer

1. EC2 / Load balancers / Create load balancer
   - application or classic
   - name: web-elb
   - http & https
   - default VPC
   - add two subnet

2. Configure security groups
   - choose "load-balancer" security group which we just setup

3. Configure routing
   - target group: web-servers-tg1
   - ping path: /ping
   - allows us to define a "healthy" web server
   - load balancer will only forward to healthy web servers

4. register targets
5. create
