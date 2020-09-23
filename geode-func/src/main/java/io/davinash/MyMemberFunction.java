package io.davinash;


import org.apache.geode.cache.RegionService;
import org.apache.geode.cache.execute.Function;
import org.apache.geode.cache.execute.FunctionContext;
import org.apache.geode.cache.execute.FunctionException;
import org.apache.geode.cache.execute.RegionFunctionContext;
import org.apache.geode.distributed.DistributedSystem;
import org.apache.geode.distributed.internal.InternalDistributedSystem;
import org.apache.geode.distributed.internal.membership.InternalDistributedMember;

/**
 * MyMemberFunction!
 */
public class MyMemberFunction implements Function {

  @Override
  public boolean hasResult() {
    return true;
  }

  @Override
  public void execute(FunctionContext fc) {
    fc.getResultSender().lastResult(fc.getMemberName());
  }

  @Override
  public String getId() {
    return "MyMemberFunction";
  }

  @Override
  public boolean optimizeForWrite() {
    return false;
  }

  @Override
  public boolean isHA() {
    return false;
  }

}
