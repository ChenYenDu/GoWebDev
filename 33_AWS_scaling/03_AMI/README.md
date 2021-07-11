# Create an AMI (Amazon Machine Image)

1. EC2 / Instances / right-click instance / create image
   - image name: web-architecture-2019-10-31
   - description: web server 2019 October 31
   - no reboot: unchecked
     - allowing your instance to reboot gives a better image

2. create image

## Launch a new instance of your AMI in a new availability zon (AZ)

1. what AZ is your current instance running in?
   - EC2 / instances / look at the availiability zone and make note of it

2. launch a new instance from your AMI
   - EC2 / AMIs / right click / launch / next: configure

3. subnet: ```<choose a different AZ>``` / next: storage / next
4. tag
   - value: web-server-0002
5. security group
   - choose the "web-tier" security group we created
6. launch
   - specify "key pair" we want the instance to use
7. launch instance

## Add new EC2 instance to load balancer's target group

1. add the new instance to the target group
2. enter load balancer DNS into a browser to see your load balancer in action
   - refresh your browser to see the switching between web-servers-sg

---

## Create auto scaling

Auto Scaling helps you maintain application availability and allows you to scale your Amazon EC2 capacity up or down automatically according to conditions you define. You can use Auto Scaling to help ensure that you are running your desired number of Amazon EC2 instances. Auto Scaling also automatically increase the number of Amazon EC2 instances during demand spikes to maintain performance and decrease capacity during lulls to redues costs. Auto Scaling is well suited both to applications that have stable demand patterns or that experience hourly, daily, or weekly variability in usage.

### Configure auto scaling

1. EC2 / autoscaling / launch configuration
2. create auto scaling group / create launch configuration
3. My AMIs / choose your AMI
   - my image name waw "web-architecture-2019-10-31"
   - next / next
4. configure details
   - name: auto-scaling-config-2019-10-31
   - next / next
5. configure security group
   - select an existing security group /select the "web-servers-sg" security group
   - next / next / create launch configuration
   - choose an existing key pair / create launch configuration

### Create auto scaling group

1. Configure auto scaling group
   - group size: this is the minimum number of instances we'll always be running
   - network: default vpc
   - subnet: choose the availability zones (AZs) into which you've launched instances
   - advanced details
      - load balancing: check "receive traffic from elastic load balancer"
      - select your load balancer
      - health check: ELB (this is what we set up)

2. Configure scaling policies
   - keep group at initial size

3. Configure tags
   - value: web-sever-auto-scaled
4. Create auto scaling group
5. Scaling policies
   - this is where we'd add polices to say when we scale up / scale down
