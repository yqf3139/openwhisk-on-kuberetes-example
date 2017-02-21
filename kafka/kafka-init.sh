for pod in `kubectl get pod -l app=kafka --no-headers | awk '{print $1}'`
do
   kubectl exec $pod -- bash -c 'unset JMX_PORT; kafka-topics.sh --create --topic completed --replication-factor 1 --partitions 1 --zookeeper zookeeper:2181' 
done
